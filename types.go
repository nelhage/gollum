package gollum

// Type is the type of a type in the language
type Type interface {
	isType()
}

// AtomicType represents a primitive atomic type
type AtomicType struct {
	Name string
}

func (a *AtomicType) isType() {}

// TypeVariable represents a type variable
type TypeVariable struct {
	// Unique counter, for debugging; TypeVariable equality is
	// computed on identity, not contents.
	Var int64

	// "" for gensym'd variables; otherwise, the name as entered
	// in the source.
	Name string
}

func (a *TypeVariable) isType() {}

// FunctionType is the type of a Function
type FunctionType struct {
	Dom   Type
	Range Type
}

func (a *FunctionType) isType() {}

// TupleType is the type of a tuple
type TupleType struct {
	Elts []Type
}

func (t *TupleType) isType() {}

// Forall is the type of a universally qualified type term
type Forall struct {
	Vars []*TypeVariable
	Type Type
}

func (f *Forall) isType() {}
