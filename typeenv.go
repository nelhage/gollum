package lambda

// TypeEnv is the type of a typing environment
type TypeEnv struct {
	Frame  map[string]Type
	Parent *TypeEnv
}

// Lookup looks up a value in an environment
func (e *TypeEnv) Lookup(name string) Type {
	if e == nil {
		return nil
	}
	if t := e.Frame[name]; t != nil {
		return t
	}
	return e.Parent.Lookup(name)
}

// Extend returns an environment that inherits from `e` but includes
// an additional set of bindings
func (e *TypeEnv) Extend(names []string, vals []Type) *TypeEnv {
	if len(names) != len(vals) {
		panic("Extend: name/value mismatch")
	}
	e = &TypeEnv{
		Frame:  make(map[string]Type, len(names)),
		Parent: e,
	}
	for i, n := range names {
		e.Frame[n] = vals[i]
	}
	return e
}
