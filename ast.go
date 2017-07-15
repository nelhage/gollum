package lambda

import "fmt"

// Pos represents a position within a file
type Pos struct {
	// Offset within the file in bytes
	Offset       uint
	Line, Column uint
}

// Loc represents a source location.
type Loc struct {
	File       string
	Begin, End Pos
}

// AST is the type of an AST in the language
type AST interface {
	Location() Loc
	isAST()
}

// Location allows Loc to be embedded anonymously into AST nodes and
// implement the Location() method of AST
func (l Loc) Location() Loc {
	return l
}

// String implements Stringer for Loc
func (l Loc) String() string {
	return fmt.Sprintf("%s:%d:%d", l.File, l.Begin.Line+1, l.Begin.Column+1)
}

// Boolean represents a boolean literal
type Boolean struct {
	Loc
	Value bool
}

func (b *Boolean) isAST() {}

// String represents a string literal
type String struct {
	Loc
	Value string
}

func (s *String) isAST() {}

// Integer represents an integer literal
type Integer struct {
	Loc
	Value int64
}

func (s *Integer) isAST() {}

// Variable represents a variable term
type Variable struct {
	Loc
	Var string
}

func (v *Variable) isAST() {}

// Abstraction represents a lambda abstraction
type Abstraction struct {
	Loc
	Vars []AST
	Body AST
}

func (a *Abstraction) isAST() {}

// TypedName represents a `var : Type` clause
type TypedName struct {
	Loc
	Name string
	Type AST
}

func (t *TypedName) isAST() {}

// Application represents a function call
type Application struct {
	Loc
	Func AST
	Args []AST
}

func (a *Application) isAST() {}

// If represents a conditional node
type If struct {
	Loc
	Condition  AST
	Consequent AST
	Alternate  AST
}

func (i *If) isAST() {}

// TyName represents a primitive type in the AST
type TyName struct {
	Loc
	Type string
}

func (t *TyName) isAST() {}

// TyTuple represents a tuple type in the AST
type TyTuple struct {
	Loc
	Elts []AST
}

func (t *TyTuple) isAST() {}

// TyArrow represents an -> type in the AST
type TyArrow struct {
	Loc
	Dom   AST
	Range AST
}

func (t *TyArrow) isAST() {}
