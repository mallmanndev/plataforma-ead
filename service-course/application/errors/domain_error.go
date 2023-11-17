package errs

import "fmt"

type DomainError struct {
	ClassName string
	Message   string
}

func NewDomainError(ClassName string, Message string) *DomainError {
	return &DomainError{
		ClassName: ClassName,
		Message:   Message,
	}
}

func (de *DomainError) Error() string {
	return fmt.Sprintf(
		"[%s] %s",
		de.ClassName,
		de.Message,
	)
}
