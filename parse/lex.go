package parse

//go:generate goyacc -o yacc.go yacc.y

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"unicode"

	"nelhage.com/lambda"
)

type lexer struct {
	off      uint
	filename string

	ioErr error
	r     offsetReader

	result lambda.AST
	error  error
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

func (l *lexer) rune() rune {
	if l.ioErr != nil {
		return 0
	}

	r, _, err := l.r.ReadRune()
	if err != nil {
		if err != io.EOF {
			l.ioErr = err
		}
		r = 0
	}
	return r
}

func (l *lexer) unread() {
	if e := l.r.UnreadRune(); e != nil {
		if l.ioErr != nil {
			l.ioErr = e
		}
	}
}

func (l *lexer) readWhile(init rune, want func(rune) bool) string {
	runes := []rune{init}
	var r rune
	for {
		r = l.rune()
		if r == 0 || !want(r) {
			break
		}
		runes = append(runes, r)
	}
	if r != 0 {
		l.unread()
	}
	return string(runes)
}

func (l *lexer) peek() rune {
	r := l.rune()
	if r != 0 {
		l.unread()
	}
	return r
}

func (l *lexer) token(t token, val interface{}) (token, interface{}, error) {
	if l.ioErr != nil {
		return 0, nil, l.ioErr
	}
	return t, val, nil
}

func (l *lexer) Loc() lambda.Loc {
	return lambda.Loc{File: l.filename, Begin: l.off, End: l.r.off}
}

func (l *lexer) next() (token, interface{}, error) {
	var r rune
	for {
		l.off = l.r.off
		r = l.rune()
		if r == 0 {
			return l.token(tokEOF, nil)
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
	if r == '-' {
		peek := l.peek()
		if r == 0 {
			return l.token(token(r), nil)
		}
		if unicode.IsNumber(peek) {
			return l.number(r)
		}
		return token(r), nil, nil
	}

	if r == '"' {
		return l.string(r)
	}
	return token(r), nil, nil
}

func (l *lexer) number(r rune) (token, interface{}, error) {
	num := l.readWhile(r, unicode.IsNumber)
	val, e := strconv.ParseInt(num, 10, 64)
	if e != nil {
		return 0, nil, e
	}
	return tokNumber, val, nil
}

func (l *lexer) ident(r rune) (token, interface{}, error) {
	word := l.readWhile(r, func(r rune) bool {
		return unicode.Is(unicode.Pc, r) ||
			unicode.IsLetter(r) ||
			unicode.IsNumber(r)
	})

	if kw := keywords[word]; kw != 0 {
		return kw, word, nil
	}
	return tokIdent, word, nil

}

// TODO: escaping
func (l *lexer) string(r rune) (token, interface{}, error) {
	word := l.readWhile(r, func(r rune) bool { return r != '"' })
	r = l.rune()
	if r != '"' {
		if l.ioErr != nil {
			return tokEOF, nil, l.ioErr
		}
		return tokEOF, nil, errors.New(`Unmatched '"'`)
	}
	return l.token(tokStr, word[1:])
}

// yacc interface

type tokenStruct struct {
	loc lambda.Loc
	val interface{}
}

func (l *lexer) Lex(lval *yySymType) int {
	tok, val, err := l.next()
	if err != nil {
		l.error = err
		return int(tokEOF)
	}

	lval.tok = &tokenStruct{l.Loc(), val}
	return int(tok)
}

func (l *lexer) Error(e string) {
	l.error = fmt.Errorf("parser error: %s", e)
}
