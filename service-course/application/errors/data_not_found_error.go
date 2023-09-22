package errs

type DataNotFoundError struct {
	message string
}

func NewDataNotFoundError(Message string) *DataNotFoundError {
	return &DataNotFoundError{Message}
}

func (e *DataNotFoundError) Error() string {
	return e.message
}
