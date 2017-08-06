package gollum

import (
	"fmt"
)

// UnboundVariable is the error of referring to an undefined variable
type UnboundVariable struct {
	Node AST
	Var  string
}

func (u UnboundVariable) Error() string {
	return fmt.Sprintf("%s: unbound variable %q",
		u.Node.Location().String(), u.Var)
}

// UnboundType is the error of referring to an undefined type name
type UnboundType struct {
	Node AST
	Name string
}

func (u UnboundType) Error() string {
	return fmt.Sprintf("%s: unknown type name %q",
		u.Node.Location().String(), u.Name)
}

// UntypedName is returned when typechecking an abstraction with an
// untyped argument
type UntypedName struct {
	Node AST
	Var  string
}

func (u UntypedName) Error() string {
	return fmt.Sprintf("%s: missing type decl for %q",
		u.Node.Location().String(), u.Var)
}

// TypeError is the type of a type error
type TypeError struct {
	Node     AST
	Got      Type
	Expected Type
}

func (t TypeError) Error() string {
	expect := PrintType(t.Expected)

	return fmt.Sprintf("%s: type error: expected %q got %q",
		t.Node.Location().String(), expect, PrintType(t.Got))
}

// OccurCheck is returned if the "occurs" check fails during
// unification
type OccurCheck struct {
	Node  AST
	Left  Type
	Right Type
}

func (o OccurCheck) Error() string {
	return fmt.Sprintf("%s: occurs check: can't construct infinite type %s = %s",
		o.Node.Location().String(),
		PrintType(o.Left),
		PrintType(o.Right),
	)
}

func bad(where string, ty Type) string {
	return fmt.Sprintf("%s: unexpected type: %#v", where, ty)
}
