package asteval

// GlobalEnv contains the default global environment
var GlobalEnv *Environment

func init() {
	globalFuncs := []struct {
		name  string
		arity int
		fn    func([]Value) (Value, error)
	}{
		{"die", 0, func([]Value) (Value, error) { panic("die") }},
		{"!", 1, func(vs []Value) (Value, error) {
			v := vs[0]
			if b := v.(*Boolean); b != nil {
				return &Boolean{!b.Val}, nil
			}
			return nil, TypeError{v, "boolean"}
		},
		},
		{"+", 2, func(vs []Value) (Value, error) {
			l := vs[0].(*Integer)
			r := vs[1].(*Integer)
			if l == nil {
				return nil, TypeError{vs[0], "integer"}
			}
			if r == nil {
				return nil, TypeError{vs[1], "integer"}
			}
			return &Integer{l.Val + r.Val}, nil
		},
		},
		{"-", 2, func(vs []Value) (Value, error) {
			l := vs[0].(*Integer)
			r := vs[1].(*Integer)
			if l == nil {
				return nil, TypeError{vs[0], "integer"}
			}
			if r == nil {
				return nil, TypeError{vs[1], "integer"}
			}
			return &Integer{l.Val - r.Val}, nil
		},
		},
		{"*", 2, func(vs []Value) (Value, error) {
			l := vs[0].(*Integer)
			r := vs[1].(*Integer)
			if l == nil {
				return nil, TypeError{vs[0], "integer"}
			}
			if r == nil {
				return nil, TypeError{vs[1], "integer"}
			}
			return &Integer{l.Val * r.Val}, nil
		},
		},
		{"dec", 1, func(vs []Value) (Value, error) {
			v := vs[0].(*Integer)
			if v == nil {
				return nil, TypeError{vs[0], "integer"}
			}
			return &Integer{v.Val - 1}, nil
		},
		},
		{"iszero", 1, func(vs []Value) (Value, error) {
			v := vs[0].(*Integer)
			if v == nil {
				return nil, TypeError{vs[0], "integer"}
			}
			return &Boolean{v.Val == 0}, nil
		},
		},
	}
	names := make([]string, len(globalFuncs))
	vals := make([]Value, len(globalFuncs))

	for i, g := range globalFuncs {
		names[i] = g.name
		vals[i] = &NativeFunction{g.arity, g.fn}
	}
	GlobalEnv = GlobalEnv.Extend(names, vals)
}
