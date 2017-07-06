package parse

import "nelhage.com/lambda"

type ParseError struct {
	Loc   lambda.Loc
	Error string
}
