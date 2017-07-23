package parse

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/nelhage/gollum"
)

var (
	// ErrExpectedProgram is the error returned by Parse when it
	// is given an input containing a type-expression at top-level
	ErrExpectedProgram = errors.New("Expected a program, but parsed a type")

	// ErrExpectedType is the error returned by Parse when it is
	// given an input containing a progream-expression at
	// top-level
	ErrExpectedType = errors.New("Expected a type, but parsed a program")
)

// Program parses an AST out of a stream. Returned locations will be
// labeled with the provided filename.
func Program(in io.Reader, filename string) (gollum.AST, error) {
	lex := &lexer{
		r: offsetReader{
			r: bufio.NewReader(in),
		},
		filename: filename,
	}
	yyErrorVerbose = true
	yyParse(lex)
	if len(lex.errors) != 0 {
		return nil, lex.errors[0]
	}
	if lex.expression == nil {
		return nil, ErrExpectedProgram
	}
	return lex.expression, nil
}

// Type parses a type expression from the input.
func Type(in io.Reader, filename string) (gollum.AST, error) {
	lex := &lexer{
		r: offsetReader{
			r: bufio.NewReader(in),
		},
		filename: filename,
		initTok:  tokIntType,
	}
	yyErrorVerbose = true
	yyParse(lex)
	if len(lex.errors) != 0 {
		return nil, lex.errors[0]
	}
	if lex.ty == nil {
		return nil, ErrExpectedType
	}
	return lex.ty, nil
}

// Error will be returned for syntax error
type Error struct {
	Loc gollum.Loc
	Err string
}

// Error implements the error interface
func (p *Error) Error() string {
	return fmt.Sprintf("%s: %s", p.Loc.String(), p.Err)
}
