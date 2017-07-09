package parse

import (
	"bufio"
	"reflect"
	"strings"
	"testing"

	"nelhage.com/lambda"
)

type fullTok struct {
	loc lambda.Loc
	tok token
	val interface{}
}

func lex(s string) []fullTok {
	r := offsetReader{r: bufio.NewReader(strings.NewReader(s))}
	l := lexer{filename: "<test>", r: r}

	var toks []fullTok
	for {
		tok, val, e := l.next()
		if e != nil {
			panic(e)
		}
		if tok == tokEOF {
			break
		}
		toks = append(toks, fullTok{
			l.Loc(),
			tok,
			val,
		})
	}
	return toks
}

func TestLex(t *testing.T) {
	l := func(b, e uint) lambda.Loc {
		return lambda.Loc{File: "<test>", Begin: b, End: e}
	}

	cases := []struct {
		str  string
		toks []fullTok
	}{
		{
			"if and 93 \ttrue",
			[]fullTok{
				{l(0, 2), tokIf, "if"},
				{l(3, 6), tokIdent, "and"},
				{l(7, 9), tokNumber, int64(93)},
				{l(11, 15), tokBoolean, "true"},
			},
		},
		{
			`"hello" "world"`,
			[]fullTok{
				{l(0, 7), tokStr, "hello"},
				{l(8, 15), tokStr, "world"},
			},
		},
		{
			`4 - -4 -x`,
			[]fullTok{
				{l(0, 1), tokNumber, int64(4)},
				{l(2, 3), token('-'), nil},
				{l(4, 6), tokNumber, int64(-4)},
				{l(7, 8), token('-'), nil},
				{l(8, 9), tokIdent, "x"},
			},
		},
	}
	for _, tc := range cases {
		toks := lex(tc.str)
		if len(toks) != len(tc.toks) {
			t.Errorf("length mismatch input=%q\nwant %#v\ngot  %#v",
				tc.str,
				tc.toks,
				toks)
		}
		for i, got := range toks {
			if i >= len(tc.toks) {
				break
			}
			if !reflect.DeepEqual(got, tc.toks[i]) {
				t.Errorf("mismatch input=%q tok=%d\nwant %#v\ngot  %#v",
					tc.str, i,
					tc.toks[i],
					got)
			}
		}
	}
}

func TestLexError(t *testing.T) {
	r := offsetReader{r: bufio.NewReader(strings.NewReader(`"foo`))}
	l := lexer{filename: "<test>", r: r}
	tok, _, e := l.next()
	if e == nil {
		t.Fatalf("did not error on mismatched \": %v", tok)
	}
}