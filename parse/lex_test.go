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
	r := locReader{r: bufio.NewReader(strings.NewReader(s))}
	l := lexer{filename: "<test>", r: r}

	var toks []fullTok
	for {
		tok, val, e := l.next()
		if e != nil {
			panic(e)
		}
		if tok == 0 {
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
	l := func(off uint) lambda.Loc {
		return lambda.Loc{File: "<test>", Char: off}
	}

	cases := []struct {
		str  string
		toks []fullTok
	}{
		{
			"if and 93 \ttrue",
			[]fullTok{
				{l(0), tokIf, "if"},
				{l(3), tokIdent, "and"},
				{l(7), tokNumber, int64(93)},
				{l(11), tokBoolean, "true"},
			},
		},
		{
			`"hello" "world"`,
			[]fullTok{
				{l(0), tokStr, "hello"},
				{l(8), tokStr, "world"},
			},
		},
		{
			`4 - -4 -x`,
			[]fullTok{
				{l(0), tokNumber, int64(4)},
				{l(2), token('-'), nil},
				{l(4), tokNumber, int64(-4)},
				{l(7), token('-'), nil},
				{l(8), tokIdent, "x"},
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
