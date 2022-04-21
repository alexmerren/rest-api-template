package entities

type baseError struct {
	Err error
	msg string
}

func (e *baseError) Error() string {
	return e.msg
}

func (e *baseError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}

func newBaseError(msg string, err error) *baseError {
	return &baseError{
		Err: err,
		msg: msg,
	}
}

type NotFoundError struct {
	*baseError
}

func NewNotFoundError(msg string, err error) *NotFoundError {
	return &NotFoundError{
		baseError: newBaseError(msg, err),
	}
}
