package puzzle

import "fmt"

// UnkownPartError occurs when given part is not defined.
type UnkownPartError struct {
	Day  Day
	Part int
}

func (e UnkownPartError) Error() string {
	return fmt.Sprintf("unknown part: %T :: %d", e.Day, e.Part)
}

// NoInputError occurs when we have no input values.
type NoInputError struct{}

func (e NoInputError) Error() string {
	return "no input, no puzzle"
}

// NoSolutionError occurs when we are unable to find a solution.
type NoSolutionError struct{}

func (e NoSolutionError) Error() string {
	return "404 - solution not found"
}

// NotImplementedError occurs when something is not implemented.
type NotImplementedError struct{}

func (e NotImplementedError) Error() string {
	return "501 - not implemented"
}

// InputParseError occurs when something is not implemented.
type InputParseError struct {
	Message string
}

func (e InputParseError) Error() string {
	return fmt.Sprintf("input parse error: %s", e.Message)
}
