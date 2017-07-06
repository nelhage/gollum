package asteval

import (
	"reflect"
	"testing"

	"nelhage.com/lambda"
)

var globalEnv *Environment

func init() {
	globalFuncs := []struct {
		name  string
		arity int
		fn    func([]Value) (Value, error)
	}{
		{"die", 0, func([]Value) (Value, error) { panic("die") }},
		{"hello", 1, func(vs []Value) (Value, error) {
			v := vs[0]
			if s := v.(*String); s != nil {
				return &String{"Hello, " + s.Val}, nil
			}
			return nil, TypeError{v, "string"}
		},
		},
		{"not", 1, func(vs []Value) (Value, error) {
			v := vs[0]
			if b := v.(*Boolean); b != nil {
				return &Boolean{!b.Val}, nil
			}
			return nil, TypeError{v, "boolean"}
		},
		},
	}
	names := make([]string, len(globalFuncs))
	vals := make([]Value, len(globalFuncs))

	for i, g := range globalFuncs {
		names[i] = g.name
		vals[i] = &NativeFunction{g.arity, g.fn}
	}
	globalEnv = globalEnv.Extend(names, vals)
}

func TestEval(t *testing.T) {
	cases := []struct {
		name string
		ast  lambda.AST
		val  Value
		err  error
	}{
		{
			"lit",
			&lambda.Boolean{Value: true},
			&Boolean{true},
			nil,
		},
		{
			"if",
			&lambda.If{
				Condition:  &lambda.Boolean{Value: true},
				Consequent: &lambda.String{Value: "true"},
				Alternate: &lambda.Application{
					Func: &lambda.Variable{Var: "die"},
					Args: []lambda.AST{&lambda.Boolean{Value: true}},
				},
			},
			&String{"true"},
			nil,
		},
		{
			"unbound",
			&lambda.Variable{Var: "foobar"},
			nil,
			UnboundVariable{"foobar"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			v, e := Eval(tc.ast, globalEnv)
			if !reflect.DeepEqual(v, tc.val) {
				t.Errorf("Bad eval: got %#v want %#v",
					v, tc.val,
				)
			}
			if !reflect.DeepEqual(e, tc.err) {
				t.Errorf("Bad err: got %#v want %#v",
					e, tc.err,
				)
			}
		})
	}
}
