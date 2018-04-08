package parse

//go:generate goyacc -o yacc.go yacc.y

import (
	"errors"
	"io"
	"strconv"
	"unicode"

	"github.com/nelhage/gollum"
)

type lexer struct {
	pos      gollum.Pos
	filename string

	ioErr error
	r     offsetReader

	expression gollum.AST
	ty         gollum.AST

	errors []error

	initTok token
}

type token int

const (
	eof = 0
)

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

func (l *lexer) Loc() gollum.Loc {
	return gollum.Loc{File: l.filename, Begin: l.pos, End: l.r.pos}
}

func (l *lexer) next() (token, interface{}, error) {
	if l.initTok != 0 {
		t := l.initTok
		l.initTok = 0
		return t, nil, nil
	}

	var r rune
skipws:
	for {
		l.pos = l.r.pos
		r = l.rune()
		switch {
		case r == 0:
			return l.token(eof, nil)
		case unicode.IsSpace(r):
			continue skipws
		case r == '#':
			l.readWhile(r, func(r rune) bool { return r != '\n' })
			l.rune()
			continue skipws
		}
		break
	}

	switch {
	case unicode.Is(unicode.Pc, r) || unicode.IsLetter(r):
		return l.ident(r)
	case unicode.IsNumber(r):
		return l.number(r)
	case r == '-':
		peek := l.peek()
		if r == 0 {
			return l.token(token(r), nil)
		}
		if unicode.IsNumber(peek) {
			return l.number(r)
		}
		if peek == '>' {
			l.rune()
			return tokArrow, nil, nil
		}
		return token(r), nil, nil
	case r == '"':
		return l.string(r)
	default:
		return token(r), nil, nil
	}
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
			return eof, nil, l.ioErr
		}
		return eof, nil, errors.New(`Unmatched '"'`)
	}
	return l.token(tokStr, word[1:])
}

// yacc interface

type tokenStruct struct {
	tok token
	loc gollum.Loc
	val interface{}
}

func (l *lexer) Lex(lval *yySymType) int {
	tok, val, err := l.next()
	if err != nil {
		l.errors = append(l.errors, err)
		return eof
	}

	lval.tok = &tokenStruct{tok, l.Loc(), val}
	return int(tok)
}

func (l *lexer) Error(e string) {
	l.errors = append(l.errors, &Error{l.Loc(), e})
}

func extend(l, r gollum.Loc) gollum.Loc {
	if l.File != r.File {
		panic("extend filename")
	}
	if l.Begin.Offset > r.End.Offset {
		panic("extend order")
	}
	return gollum.Loc{
		File:  l.File,
		Begin: l.Begin,
		End:   r.End,
	}
}

func arithmetic(op *tokenStruct, args ...gollum.AST) gollum.AST {
	return &gollum.Application{
		Loc: extend(args[0].Location(), args[len(args)-1].Location()),
		Func: &gollum.Variable{
			Loc: op.loc,
			Var: string(rune(op.tok)),
		},
		Args: args,
	}
}
