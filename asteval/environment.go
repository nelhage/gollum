package asteval

// Environment is the type of a runtime environment
type Environment struct {
	Var    string
	Value  Value
	Parent *Environment
}

// Lookup looks up a value in an environment
func (e *Environment) Lookup(name string) Value {
	if e == nil {
		return nil
	}
	if name == e.Var {
		return e.Value
	}
	return e.Parent.Lookup(name)
}

// Extend returns an environment that inherits from `e` but includes
// an additional binding
func (e *Environment) Extend(name string, val Value) *Environment {
	return &Environment{
		Var:    name,
		Value:  val,
		Parent: e,
	}
}
