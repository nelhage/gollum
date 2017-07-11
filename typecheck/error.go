package typecheck

import (
	"fmt"
	"strings"

	"nelhage.com/lambda"
)

// UnboundVariable is the error of referring to an undefined variable
type UnboundVariable struct {
	Node lambda.AST
	Var  string
}

func (u UnboundVariable) Error() string {
	return fmt.Sprintf("%s: unbound variable %q",
		u.Node.Location().String(), u.Var)
}

// UnboundType is the error of referring to an undefined type name
type UnboundType struct {
	Node lambda.AST
	Name string
}

func (u UnboundType) Error() string {
	return fmt.Sprintf("%s: unknown type name %q",
		u.Node.Location().String(), u.Name)
}

// UntypedName is returned when typechecking an abstraction with an
// untyped argument
type UntypedName struct {
	Node lambda.AST
	Var  string
}

func (u UntypedName) Error() string {
	return fmt.Sprintf("%s: missing type decl for %q",
		u.Node.Location().String(), u.Var)
}

// TypeError is the type of a type error
type TypeError struct {
	Node       lambda.AST
	Got        lambda.Type
	Expected   string
	ExpectedTy lambda.Type
}

func (t TypeError) Error() string {
	var expect string
	if t.ExpectedTy == nil {
		expect = t.Expected
	} else {
		expect = PrintType(t.ExpectedTy)
	}

	return fmt.Sprintf("%s: type error: expected %q got %q",
		t.Node.Location().String(), expect, PrintType(t.Got))
}

// PrintType returns a string representation of a type
func PrintType(t lambda.Type) string {
	switch n := t.(type) {
	case *lambda.AtomicType:
		return n.Name
	case *lambda.FunctionType:
		r := PrintType(n.Range)
		if _, ok := n.Range.(*lambda.FunctionType); ok {
			r = fmt.Sprintf("(%s)", r)
		}
		return fmt.Sprintf("%s -> %s", PrintType(n.Dom), r)
	case *lambda.TupleType:
		var bits []string
		for _, e := range n.Elts {
			bits = append(bits, PrintType(e))
		}
		return fmt.Sprintf("(%s,)", strings.Join(bits, ", "))
	default:
		panic(fmt.Sprintf("unknown type: %#v", t))
	}
}
