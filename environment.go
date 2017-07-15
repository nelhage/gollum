package lambda

// Environment is the type of a typing environment
type Environment struct {
	Frame  map[string]Type
	Parent *Environment
}

// Lookup looks up a value in an environment
func (e *Environment) Lookup(name string) Type {
	if e == nil {
		return nil
	}
	if t := e.Frame[name]; t != nil {
		return t
	}
	return e.Parent.Lookup(name)
}

// Extend returns an environment that inherits from `e` but includes
// an additional binding
func (e *Environment) Extend(names []string, vals []Type) *Environment {
	if len(names) != len(vals) {
		panic("Extend: name/value mismatch")
	}
	e = &Environment{
		Frame:  make(map[string]Type, len(names)),
		Parent: e,
	}
	for i, n := range names {
		e.Frame[n] = vals[i]
	}
	return e
}
