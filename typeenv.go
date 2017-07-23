package gollum

import "fmt"

// TypeEnv is the type of a typing environment
type TypeEnv struct {
	Frame  map[string]Type
	Types  map[string]Type
	Parent *TypeEnv
	Vars   []*TypeVariable
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

// LookupType looks up a type name in an environment
func (e *TypeEnv) LookupType(name string) Type {
	return e.Types[name]
}

// NewVar instantiates a new type variable, and stores it in the frame
func (e *TypeEnv) NewVar(name string) *TypeVariable {
	if t := e.Lookup(name); t != nil {
		panic(fmt.Sprintf("NewVar: redefining type %q", name))
	}

	tyvar := &TypeVariable{Name: name}
	e.Types[name] = tyvar
	return tyvar
}

// Extend returns an environment that inherits from `e` but includes
// an additional set of bindings
func (e *TypeEnv) Extend(names []string, vals []Type, vars []*TypeVariable) *TypeEnv {
	if len(names) != len(vals) {
		panic("Extend: name/value mismatch")
	}
	var types map[string]Type
	if e != nil {
		types = e.Types
	}
	e = &TypeEnv{
		Frame:  make(map[string]Type, len(names)),
		Parent: e,
		Vars:   vars,
		Types:  types,
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
func (e *TypeEnv) BoundVars() []*TypeVariable {
	var out []*TypeVariable
	for e != nil {
		out = append(out, e.Vars...)
		e = e.Parent
	}
	return out
}
