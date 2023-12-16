package statuscode

type Code interface {
	Message() string
	Code() int
}
