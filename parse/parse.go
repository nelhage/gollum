package parse

import (
	"bufio"
	"fmt"
	"github.com/nelhage/gollum"
	"io"
)

// Parse parses an AST out of a stream. Returned locations will be
// labeled with the provided filename.
func Parse(in io.Reader, filename string) (gollum.AST, error) {
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
	return lex.result, nil
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
