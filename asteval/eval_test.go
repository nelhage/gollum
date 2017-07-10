package asteval

import (
	"reflect"
	"testing"

	"nelhage.com/lambda"
)

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
			v, e := Eval(tc.ast, GlobalEnv)
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
