package errors

type (
	ErrNotFound struct {
		msg string
	}

	ErrBadRequest struct {
		msg string
	}
)

func (err ErrNotFound) Error() string {
	return err.msg
}

func NewErrNotFound(msg string) ErrNotFound {
	return ErrNotFound{msg: msg}
}

func (err ErrBadRequest) Error() string {
	return err.msg
}

func NewErrBadRequest(msg string) ErrBadRequest {
	return ErrBadRequest{msg: msg}
}
