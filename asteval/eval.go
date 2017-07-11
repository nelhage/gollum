package asteval

import (
	"fmt"

	"nelhage.com/lambda"
)

// Eval evaluate an AST node within the specified Environment
func Eval(a lambda.AST, e *Environment) (Value, error) {
	switch n := a.(type) {
	case *lambda.Boolean:
		return &Boolean{n.Value}, nil
	case *lambda.String:
		return &String{n.Value}, nil
	case *lambda.Integer:
		return &Integer{n.Value}, nil
	case *lambda.Variable:
		v := e.Lookup(n.Var)
		if v == nil {
			return nil, UnboundVariable{n.Var}
		}
		return v, nil
	case *lambda.Abstraction:
		var names []string
		for _, v := range n.Vars {
			names = append(names, v.(*lambda.TypedName).Name)
		}
		return &Closure{
			Env:  e,
			Args: names,
			Body: n.Body,
		}, nil
	case *lambda.Application:
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
	case *lambda.If:
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
