package parse

import (
	"errors"
	"io"

	"nelhage.com/lambda"
)

// Parse parses an AST out of a stream. Returned locations will be
// labeled with the provided filename.
func Parse(in io.Reader, filename string) (lambda.AST, error) {
	return nil, errors.New("unimplemented")
}
