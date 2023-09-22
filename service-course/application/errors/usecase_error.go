package errs

import "fmt"

type UseCaseError struct {
	usecase   string
	message   string
	wrapError error
}

func (uc *UseCaseError) Error() string {
	if uc.wrapError != nil {
		return fmt.Sprintf("[%s] %s: %s", uc.usecase, uc.message, uc.wrapError)
	}
	return fmt.Sprintf("[%s] %s.", uc.usecase, uc.message)
}

func (uc *UseCaseError) Unwrap() error {
	return uc.wrapError
}

func NewCreateUserUseCaseError(Message string, error error) *UseCaseError {
	return &UseCaseError{
		usecase:   "Create User",
		message:   Message,
		wrapError: error,
	}
}

func NewUpdateCourseUseCaseError(Message string, error error) *UseCaseError {
	return &UseCaseError{
		usecase:   "Update Course",
		message:   Message,
		wrapError: error,
	}
}

func NewDeleteCourseUseCaseError(Message string, error error) *UseCaseError {
	return &UseCaseError{
		usecase:   "Delete Course",
		message:   Message,
		wrapError: error,
	}
}

func NewCreateSectionUseCaseError(Message string, error error) *UseCaseError {
	return &UseCaseError{
		usecase:   "Create Section",
		message:   Message,
		wrapError: error,
	}
}

func NewUpdateSectionUseCaseError(Message string, error error) *UseCaseError {
	return &UseCaseError{
		usecase:   "Update Section",
		message:   Message,
		wrapError: error,
	}
}

func NewDeleteSectionUseCaseError(Message string, error error) *UseCaseError {
	return &UseCaseError{
		usecase:   "Delete Section",
		message:   Message,
		wrapError: error,
	}
}
