package asteval

import "fmt"

// Environment is the type of a runtime environment
type Environment struct {
	Frame  map[string]Value
	Parent *Environment
}

// Lookup looks up a value in an environment
func (e *Environment) Lookup(name string) Value {
	if e == nil {
		return nil
	}
	if v := e.Frame[name]; v != nil {
		return v
	}
	return e.Parent.Lookup(name)
}

// Extend returns an environment that inherits from `e` but includes
// an additional binding
func (e *Environment) Extend(names []string, vals []Value) *Environment {
	if len(names) != len(vals) {
		panic("Extend: name/value mismatch")
	}
	e = &Environment{
		Frame:  make(map[string]Value, len(names)),
		Parent: e,
	}
	for i, n := range names {
		e.Frame[n] = vals[i]
	}
	return e
}

// SetLocal sets a number of name, value pairs in the local frame
func (e *Environment) SetLocal(names []string, vals []Value) {
	for i, n := range names {
		if _, ok := e.Frame[n]; !ok {
			panic(fmt.Sprintf("SetLocal: %s", n))
		}
		e.Frame[n] = vals[i]
	}
}
