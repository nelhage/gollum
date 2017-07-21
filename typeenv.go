package gollum

import "fmt"

// TypeEnv is the type of a typing environment
type TypeEnv struct {
	Frame  map[string]Type
	Parent *TypeEnv
	Vars   []int64
}

// Lookup looks up a value in an environment
func (e *TypeEnv) Lookup(name string) Type {
	for e != nil {
		if t := e.Frame[name]; t != nil {
			return t
		}
		e = e.Parent
	}

	return nil
}

// Extend returns an environment that inherits from `e` but includes
// an additional set of bindings
func (e *TypeEnv) Extend(names []string, vals []Type, vars []int64) *TypeEnv {
	if len(names) != len(vals) {
		panic("Extend: name/value mismatch")
	}
	e = &TypeEnv{
		Frame:  make(map[string]Type, len(names)),
		Parent: e,
		Vars:   vars,
	}
	for i, n := range names {
		e.Frame[n] = vals[i]
	}
	return e
}

// SetLocal sets a number of name, value pairs in the local frame
func (e *TypeEnv) SetLocal(names []string, vals []Type) {
	if len(names) != len(vals) {
		panic("SetLocal: name/value mismatch")
	}
	for i, n := range names {
		if _, ok := e.Frame[n]; !ok {
			panic(fmt.Sprintf("SetLocal: %s", n))
		}
		e.Frame[n] = vals[i]
	}
}

// BoundVars returns a list of all type variables bound in e
func (e *TypeEnv) BoundVars() []int64 {
	var out []int64
	for e != nil {
		out = append(out, e.Vars...)
		e = e.Parent
	}
	return out
}
