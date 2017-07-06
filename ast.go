package lambda

// Loc represents a source location.
type Loc struct {
	File string
	Char uint
}

// AST is the type of an AST in the language
type AST interface {
	Loc() Loc
	isAST()
}

type hasLoc struct {
	loc Loc
}

func (loc *hasLoc) Loc() Loc {
	return Loc{}
}

// Boolean represents a boolean literal
type Boolean struct {
	hasLoc
	Value bool
}

func (b *Boolean) isAST() {}

// String represents a string literal
type String struct {
	hasLoc
	Value string
}

func (s *String) isAST() {}

// Integer represents an integer literal
type Integer struct {
	hasLoc
	Value int64
}

func (s *Integer) isAST() {}

// Variable represents a variable term
type Variable struct {
	hasLoc
	Var string
}

func (v *Variable) isAST() {}

// Abstraction represents a lambda abstraction
type Abstraction struct {
	hasLoc
	Vars []string
	Body AST
}

func (a *Abstraction) isAST() {}

// Application represents a function call
type Application struct {
	hasLoc
	Func AST
	Args []AST
}

func (a *Application) isAST() {}

// If represents a conditional node
type If struct {
	hasLoc
	Condition  AST
	Consequent AST
	Alternate  AST
}

func (i *If) isAST() {}
