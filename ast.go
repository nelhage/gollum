package lambda

// AST is the type of an AST in the language
type AST interface {
	isAST()
}

// Boolean represents a boolean literal
type Boolean struct {
	Value bool
}

func (b *Boolean) isAST() {}

// String represents a string literal
type String struct {
	Value string
}

func (s *String) isAST() {}

// Variable represents a variable term
type Variable struct {
	Var string
}

func (v *Variable) isAST() {}

// Abstraction represents a lambda abstraction
type Abstraction struct {
	Var  string
	Body AST
}

func (a *Abstraction) isAST() {}

// Application represents a function call
type Application struct {
	Func AST
	Arg  AST
}

func (a *Application) isAST() {}

// If represents a conditional node
type If struct {
	Condition  AST
	Consequent AST
	Alternate  AST
}

func (i *If) isAST() {}
