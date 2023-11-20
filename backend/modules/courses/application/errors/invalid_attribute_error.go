package errs

import "fmt"

type InvalidAttributeError struct {
	ClassName     string
	AttributeName string
	Reason        string
}

func NewInvalidAttributeError(ClassName string, AttributeName string, Reason string) *InvalidAttributeError {
	return &InvalidAttributeError{
		ClassName:     ClassName,
		AttributeName: AttributeName,
		Reason:        Reason,
	}
}

func (ia *InvalidAttributeError) Error() string {
	return fmt.Sprintf("[%s] Invalid '%s': %s.",
		ia.ClassName,
		ia.AttributeName,
		ia.Reason)
}
