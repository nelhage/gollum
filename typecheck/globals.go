package typecheck

import "nelhage.com/lambda"

// GlobalEnv contains the default global environment
var GlobalEnv *Environment

// GlobalTypes stores a lookup table for global type names
var GlobalTypes map[string]lambda.Type

var (
	boolType = &lambda.AtomicType{Name: "bool"}
	intType  = &lambda.AtomicType{Name: "int"}
	strType  = &lambda.AtomicType{Name: "str"}
)

func init() {
	var binaryIntFn = &lambda.FunctionType{
		Dom: &lambda.TupleType{
			Elts: []lambda.Type{
				intType, intType,
			},
		},
		Range: intType,
	}

	globalFuncs := []struct {
		name string
		ty   lambda.Type
	}{
		{"die", &lambda.TupleType{}},
		{"not", &lambda.FunctionType{
			Dom: &lambda.TupleType{
				Elts: []lambda.Type{boolType},
			},
			Range: boolType,
		}},
		{"add", binaryIntFn},
		{"sub", binaryIntFn},
		{"mul", binaryIntFn},
		{"dec", &lambda.FunctionType{
			Dom: &lambda.TupleType{
				Elts: []lambda.Type{intType},
			},
			Range: intType,
		}},
		{"iszero", &lambda.FunctionType{
			Dom: &lambda.TupleType{
				Elts: []lambda.Type{intType},
			},
			Range: boolType,
		}},
	}
	names := make([]string, len(globalFuncs))
	types := make([]lambda.Type, len(globalFuncs))

	for i, g := range globalFuncs {
		names[i] = g.name
		types[i] = g.ty
	}
	GlobalEnv = GlobalEnv.Extend(names, types)

	GlobalTypes = make(map[string]lambda.Type)

	GlobalTypes[intType.Name] = intType
	GlobalTypes[boolType.Name] = boolType
	GlobalTypes[strType.Name] = strType
}
