package gollum

import (
	"bytes"
	"fmt"
	"strings"
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

// PrintType returns a string representation of a type
func PrintType(t Type) string {
	switch n := t.(type) {
	case *AtomicType:
		return n.Name
	case *FunctionType:
		dom := n.Dom
		var d string
		if dtup, ok := n.Dom.(*TupleType); ok && len(dtup.Elts) == 1 {
			dom = dtup.Elts[0]
		}

		if _, ok := dom.(*FunctionType); ok {
			d = fmt.Sprintf("(%s)", PrintType(dom))
		} else {
			d = PrintType(dom)
		}
		r := PrintType(n.Range)
		return fmt.Sprintf("%s -> %s", d, r)
	case *TupleType:
		var bits []string
		for _, e := range n.Elts {
			bits = append(bits, PrintType(e))
		}
		return fmt.Sprintf("(%s)", strings.Join(bits, ", "))
	case *TypeVariable:
		return varname(n.Var)
	case *Forall:
		var b bytes.Buffer
		b.WriteString("∀")
		for i, v := range n.Vars {
			b.WriteString(varname(v))
			if i != len(n.Vars)-1 {
				b.WriteString(",")
			}
		}
		b.WriteString(".")
		b.WriteString(PrintType(n.Type))
		return b.String()
	default:
		panic(fmt.Sprintf("unknown type: %#v", t))
	}
}

func bad(where string, ty Type) string {
	return fmt.Sprintf("%s: unexpected type: %#v", where, ty)
}

func varname(v int64) string {
	// ceil(log_26(2**64))
	var buf [14]byte
	i := len(buf)

	for {
		i--
		buf[i] = byte('A' + (v % 26))
		v /= 26
		if v == 0 {
			break
		}
	}

	return string(buf[i:])
}
