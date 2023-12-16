package statuscode

const (
	CodeSuccess = iota
	CodeUnknownError

	// TODO add your code
)

var (
	Success      = NewSuccess(CodeSuccess, "success")
	UnknownError = NewError(CodeUnknownError, "unknown error")
)
