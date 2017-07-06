package parse

import (
	"io"
	"strconv"
	"unicode"

	"nelhage.com/lambda"
)

type lexer struct {
	off      uint
	filename string

	r locReader
}

type token int

const (
	tokEOF token = iota

	// identifiers
	tokFunc
	tokIf

	// primitives
	tokBoolean
	tokNumber
	tokStr

	// identifiers
	tokIdent
)

var keywords = map[string]token{
	"if":    tokIf,
	"fn":    tokFunc,
	"true":  tokBoolean,
	"false": tokBoolean,
}

func (l *lexer) readWhile(init rune, want func(rune) bool) (string, error) {
	runes := []rune{init}
	var r rune
	var e error
	for {
		r, _, e = l.r.ReadRune()
		if e == io.EOF {
			break
		}
		if e != nil {
			return "", e
		}
		if !want(r) {
			break
		}
		runes = append(runes, r)
	}
	if r != 0 {
		if e := l.r.UnreadRune(); e != nil {
			return "", e
		}
	}
	return string(runes), nil
}

func (l *lexer) Loc() lambda.Loc {
	return lambda.Loc{File: l.filename, Char: l.off}
}

func (l *lexer) next() (token, interface{}, error) {
	var r rune
	var e error
	for {
		l.off = l.r.off
		r, _, e = l.r.ReadRune()
		if e == io.EOF {
			return tokEOF, nil, nil
		}

		if e != nil {
			return 0, nil, e
		}
		if !unicode.IsSpace(r) {
			break
		}
	}

	if unicode.Is(unicode.Pc, r) || unicode.IsLetter(r) {
		return l.ident(r)
	}
	if unicode.IsNumber(r) {
		return l.number(r)
	}
	if r == '"' {
		return l.string(r)
	}
	return token(r), nil, nil
}

func (l *lexer) number(r rune) (token, interface{}, error) {
	num, e := l.readWhile(r, unicode.IsNumber)
	if e != nil {
		return 0, nil, e
	}
	val, e := strconv.ParseInt(num, 10, 64)
	if e != nil {
		return 0, nil, e
	}
	return tokNumber, val, nil
}

func (l *lexer) ident(r rune) (token, interface{}, error) {
	word, e := l.readWhile(r, func(r rune) bool {
		return unicode.Is(unicode.Pc, r) ||
			unicode.IsLetter(r) ||
			unicode.IsNumber(r)
	})
	if e != nil {
		return 0, nil, e
	}
	if kw := keywords[word]; kw != 0 {
		return kw, word, nil
	}
	return tokIdent, word, nil

}

func (l *lexer) string(r rune) (token, interface{}, error) {
	word, e := l.readWhile(r, func(r rune) bool { return r != '"' })
	if e != nil {
		return 0, nil, e
	}
	l.r.ReadRune()
	return tokStr, word[1:], nil
}
