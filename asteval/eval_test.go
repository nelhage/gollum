package asteval

import (
	"reflect"
	"testing"

	"nelhage.com/lambda"
)

var globalEnv *Environment

func init() {
	globalFuncs := []struct {
		name string
		fn   func([]Value) (Value, error)
	}{
		// TODO: arity-checking
		{"die", func([]Value) (Value, error) { panic("die") }},
		{"hello", func(vs []Value) (Value, error) {
			v := vs[0]
			if s := v.(*String); s != nil {
				return &String{"Hello, " + s.Val}, nil
			}
			return nil, TypeError{v, "string"}
		},
		},
		{"not", func(vs []Value) (Value, error) {
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
		vals[i] = &NativeFunction{g.fn}
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
			&lambda.Boolean{true},
			&Boolean{true},
			nil,
		},
		{
			"if",
			&lambda.If{
				Condition:  &lambda.Boolean{true},
				Consequent: &lambda.String{"true"},
				Alternate: &lambda.Application{
					Func: &lambda.Variable{"die"},
					Args: []lambda.AST{&lambda.Boolean{true}},
				},
			},
			&String{"true"},
			nil,
		},
		{
			"unbound",
			&lambda.Variable{"foobar"},
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
