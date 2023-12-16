package statuscode

type SuccessCode struct {
	msg  string
	code int
}

func NewSuccess(code int, msg string) *SuccessCode {
	return &SuccessCode{
		msg:  msg,
		code: code,
	}
}

func (s *SuccessCode) Message() string {
	return s.msg
}

func (s *SuccessCode) Code() int {
	return s.code
}

func (s *SuccessCode) Clone(msg string) *SuccessCode {
	return &SuccessCode{
		code: s.code,
		msg:  msg,
	}
}
