package asteval

import lambda "github.com/nelhage/gollum"

// Value is a runtime value
type Value interface {
	isValue()
}

// Boolean is a boolean value
type Boolean struct {
	Val bool
}

func (b *Boolean) isValue() {}

// String is a string value
type String struct {
	Val string
}

func (s *String) isValue() {}

// Integer is an integer value
type Integer struct {
	Val int64
}

func (s *Integer) isValue() {}

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
