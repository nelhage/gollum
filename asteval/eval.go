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
	case *lambda.Variable:
		v := e.Lookup(n.Var)
		if v == nil {
			return nil, UnboundVariable{n.Var}
		}
		return v, nil
	case *lambda.Abstraction:
		return &Closure{
			Env:  e,
			Arg:  n.Var,
			Body: n.Body,
		}, nil
	case *lambda.Application:
		fn, err := Eval(n.Func, e)
		if err != nil {
			return nil, err
		}
		arg, err := Eval(n.Arg, e)
		if err != nil {
			return nil, err
		}
		return evalFunc(fn, arg, e)
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

func evalFunc(fn Value, arg Value, e *Environment) (Value, error) {
	switch f := fn.(type) {
	case *Closure:
		e := f.Env.Extend(f.Arg, arg)
		return Eval(f.Body, e)
	case *NativeFunction:
		return f.Func(arg)
	default:
		return nil, TypeError{fn, "function"}
	}
}
