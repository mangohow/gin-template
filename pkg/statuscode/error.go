package statuscode

import "fmt"

type ErrorCode struct {
	err  error
	msg  string
	code int
}

func NewError(code int, msg string) *ErrorCode {
	return &ErrorCode{
		err:  nil,
		msg:  msg,
		code: code,
	}
}

func NewWithError(err error, code int, msg string) *ErrorCode {
	return &ErrorCode{
		err:  err,
		msg:  msg,
		code: code,
	}
}

func (e *ErrorCode) Wrap(err error) {
	e.err = err
}

func (e *ErrorCode) Error() string {
	if e.err != nil {
		return fmt.Sprintf("statuscode: %d, msg: %s, err: %s", e.code, e.msg, e.err.Error())
	}
	return fmt.Sprintf("statuscode: %d, msg: %s", e.code, e.msg)
}

func (e *ErrorCode) Code() int {
	return e.code
}

func (e *ErrorCode) Message() string {
	return e.msg
}

func (e *ErrorCode) Clone(err error) *ErrorCode {
	return &ErrorCode{
		err:  err,
		msg:  e.msg,
		code: e.code,
	}
}
