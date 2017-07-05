package asteval

import "nelhage.com/lambda"

// Value is a runtime value
type Value interface {
	isValue()
}

// Boolean is a boolean literal
type Boolean struct {
	Val bool
}

func (b *Boolean) isValue() {}

// String is a string literal
type String struct {
	Val string
}

func (s *String) isValue() {}

// NativeFunction is a built-in implemented in the runtime
type NativeFunction struct {
	Arity int
	Func  func([]Value) (Value, error)
}

func (n *NativeFunction) isValue() {}

// Closure is an in-language function
type Closure struct {
	Env  *Environment
	Args []string
	Body lambda.AST
}

func (c *Closure) isValue() {}
