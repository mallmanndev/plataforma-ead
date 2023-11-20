package errs

import "fmt"

type NotFoundError struct {
	name string
}

func NewNotFoundError(Name string) *NotFoundError {
	return &NotFoundError{Name}
}

func (nf *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found.", nf.name)
}
