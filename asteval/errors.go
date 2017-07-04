package asteval

import "fmt"

// UnboundVariable is the error returned by referencing an unbound
// variable
type UnboundVariable struct {
	Variable string
}

func (u UnboundVariable) Error() string {
	return fmt.Sprintf("Unbound variable: %s", u.Variable)
}

// TypeError is a runtime type error
type TypeError struct {
	Value    Value
	Expected string
}

func (t TypeError) Error() string {
	return fmt.Sprintf("Expected type: %q", t.Expected)
}
