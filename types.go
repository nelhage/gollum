package lambda

// Type is the type of a type in the language
type Type interface {
	isType()
}

// AtomicType represents a primitive atomic type
type AtomicType struct {
	TyCon string
}

func (a *AtomicType) isType() {}

// FunctionType is the type of a Function
type FunctionType struct {
	Dom   Type
	Range Type
}

func (a *FunctionType) isType() {}
