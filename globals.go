package lambda

// GlobalEnv contains the default global environment
var GlobalEnv *Environment

// GlobalTypes stores a lookup table for global type names
var GlobalTypes map[string]Type

var (
	boolType = &AtomicType{Name: "bool"}
	intType  = &AtomicType{Name: "int"}
	strType  = &AtomicType{Name: "str"}
	unitType = &TupleType{}
)

func init() {
	var binaryIntFn = &FunctionType{
		Dom: &TupleType{
			Elts: []Type{
				intType, intType,
			},
		},
		Range: intType,
	}

	globalFuncs := []struct {
		name string
		ty   Type
	}{
		{"die", &FunctionType{
			Dom:   unitType,
			Range: unitType,
		}},
		{"not", &FunctionType{
			Dom: &TupleType{
				Elts: []Type{boolType},
			},
			Range: boolType,
		}},
		{"add", binaryIntFn},
		{"sub", binaryIntFn},
		{"mul", binaryIntFn},
		{"dec", &FunctionType{
			Dom: &TupleType{
				Elts: []Type{intType},
			},
			Range: intType,
		}},
		{"iszero", &FunctionType{
			Dom: &TupleType{
				Elts: []Type{intType},
			},
			Range: boolType,
		}},
	}
	names := make([]string, len(globalFuncs))
	types := make([]Type, len(globalFuncs))

	for i, g := range globalFuncs {
		names[i] = g.name
		types[i] = g.ty
	}
	GlobalEnv = GlobalEnv.Extend(names, types)

	GlobalTypes = make(map[string]Type)

	GlobalTypes[intType.Name] = intType
	GlobalTypes[boolType.Name] = boolType
	GlobalTypes[strType.Name] = strType
}
