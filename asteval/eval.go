package asteval

import (
	"fmt"
	"github.com/nelhage/gollum"
)

// Eval evaluate an AST node within the specified Environment
func Eval(a gollum.AST, e *Environment) (Value, error) {
	switch n := a.(type) {
	case *gollum.Boolean:
		return &Boolean{n.Value}, nil
	case *gollum.String:
		return &String{n.Value}, nil
	case *gollum.Integer:
		return &Integer{n.Value}, nil
	case *gollum.Variable:
		v := e.Lookup(n.Var)
		if v == nil {
			return nil, UnboundVariable{n.Var}
		}
		return v, nil
	case *gollum.Abstraction:
		var names []string
		for _, v := range n.Vars {
			names = append(names, v.(*gollum.TypedName).Name)
		}
		return &Closure{
			Env:  e,
			Args: names,
			Body: n.Body,
		}, nil
	case *gollum.Application:
		fn, err := Eval(n.Func, e)
		if err != nil {
			return nil, err
		}
		argv := make([]Value, len(n.Args))
		for i, ast := range n.Args {
			argv[i], err = Eval(ast, e)
			if err != nil {
				return nil, err
			}
		}
		if err != nil {
			return nil, err
		}
		return evalFunc(fn, argv, e)
	case *gollum.If:
		cond, err := Eval(n.Condition, e)
		if err != nil {
			return nil, err
		}
		b := cond.(*Boolean)
		if b == nil {
			return nil, TypeError{cond, "boolean"}
		}

		if b.Val {
			return Eval(n.Consequent, e)
		}
		return Eval(n.Alternate, e)

	case *gollum.Let:
		var names []string
		var vals []Value
		for _, v := range n.Bindings {
			nb := v.(*gollum.NameBinding)
			name := nb.Var.(*gollum.TypedName).Name
			names = append(names, name)
			vals = append(vals, &Integer{})
		}
		if n.Recursive {
			e = e.Extend(names, vals)
		}
		for i, v := range n.Bindings {
			nb := v.(*gollum.NameBinding)
			val, err := Eval(nb.Value, e)
			if err != nil {
				return nil, err
			}
			vals[i] = val
		}
		if n.Recursive {
			e.SetLocal(names, vals)
		} else {
			e = e.Extend(names, vals)
		}
		return Eval(n.Body, e)

	default:
		panic(fmt.Sprintf("unknown ast: %#v", a))
	}
}

func evalFunc(fn Value, argv []Value, e *Environment) (Value, error) {
	switch f := fn.(type) {
	case *Closure:
		if len(f.Args) != len(argv) {
			return nil, ArityError{fn, len(f.Args), argv}
		}
		e := f.Env.Extend(f.Args, argv)
		return Eval(f.Body, e)
	case *NativeFunction:
		if f.Arity >= 0 {
			if len(argv) != f.Arity {
				return nil, ArityError{fn, f.Arity, argv}
			}
		}
		return f.Func(argv)
	default:
		return nil, TypeError{fn, "function"}
	}
}
