package lambda

// Loc represents a source location.
type Loc struct {
	File       string
	Begin, End uint
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
	Vars []string
	Body AST
}

func (a *Abstraction) isAST() {}

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
