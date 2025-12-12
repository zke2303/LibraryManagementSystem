package errcode

type Error struct {
	Code int
	Msg  string
	Err  error
}

func (e *Error) Error() string {
	return e.Msg
}

func NewError(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func (err *Error) Wrap(errors error) *Error {
	return &Error{
		Code: err.Code,
		Msg:  err.Msg,
		Err:  errors,
	}
}

func (err *Error) UnWrap() error {
	return err.Err
}
