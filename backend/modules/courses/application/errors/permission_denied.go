package errs

import "fmt"

type PermissionDeniedError struct {
	action string
}

func NewPermissionDeniedError(Action string) *PermissionDeniedError {
	return &PermissionDeniedError{
		action: Action,
	}
}

func (pd *PermissionDeniedError) Error() string {
	return fmt.Sprintf("Permission denied to %s.", pd.action)
}
