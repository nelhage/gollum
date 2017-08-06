package gollum_test

import (
	"testing"

	"github.com/nelhage/gollum"
)

func TestPrintType(t *testing.T) {
	intType := &gollum.AtomicType{Name: "int"}
	var21 := &gollum.TypeVariable{Var: 21}

	cases := []struct {
		in  gollum.Type
		out string
	}{
		{intType, "int"},
		{&gollum.TypeVariable{Name: "a"}, "a"},
		{&gollum.TypeVariable{Var: 0}, "a"},
		{&gollum.FunctionType{
			Dom: &gollum.TupleType{
				Elts: []gollum.Type{intType},
			},
			Range: intType,
		}, "int -> int"},
		{&gollum.FunctionType{
			Dom: &gollum.TupleType{
				Elts: []gollum.Type{intType},
			},
			Range: &gollum.FunctionType{
				Dom: &gollum.TupleType{
					Elts: []gollum.Type{intType},
				},
				Range: intType,
			},
		}, "int -> int -> int"},
		{&gollum.FunctionType{
			Dom: &gollum.TupleType{
				Elts: []gollum.Type{
					&gollum.FunctionType{
						Dom: &gollum.TupleType{
							Elts: []gollum.Type{intType},
						},
						Range: intType,
					},
				},
			},
			Range: intType,
		}, "(int -> int) -> int"},
		{&gollum.FunctionType{
			Dom: &gollum.TupleType{
				Elts: []gollum.Type{intType, &gollum.TypeVariable{Name: "b"}},
			},
			Range: intType,
		}, "(int, b) -> int"},
		{&gollum.FunctionType{
			Dom: &gollum.TupleType{
				Elts: []gollum.Type{
					&gollum.TypeVariable{Name: "a"},
					&gollum.TypeVariable{Var: 13},
					&gollum.TypeVariable{Name: "c"},
					var21,
				},
			},
			Range: var21,
		}, "(a, b, c, d) -> d"},
		{&gollum.Forall{
			Vars: []*gollum.TypeVariable{var21},
			Type: var21,
		}, "∀a.a"},
	}

	for i, tc := range cases {
		got := gollum.PrintType(tc.in)
		if got != tc.out {
			t.Errorf("[%d] want=%q got=%q",
				i, tc.out, got,
			)
		}
	}
}
