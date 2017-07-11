package typecheck

import "nelhage.com/lambda"

// Environment is the type of a typing environment
type Environment struct {
	Frame  map[string]lambda.Type
	Parent *Environment
}

// Lookup looks up a value in an environment
func (e *Environment) Lookup(name string) lambda.Type {
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
func (e *Environment) Extend(names []string, vals []lambda.Type) *Environment {
	if len(names) != len(vals) {
		panic("Extend: name/value mismatch")
	}
	e = &Environment{
		Frame:  make(map[string]lambda.Type, len(names)),
		Parent: e,
	}
	for i, n := range names {
		e.Frame[n] = vals[i]
	}
	return e
}
