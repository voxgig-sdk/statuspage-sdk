package core

type StatuspageError struct {
	IsStatuspageError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewStatuspageError(code string, msg string, ctx *Context) *StatuspageError {
	return &StatuspageError{
		IsStatuspageError: true,
		Sdk:              "Statuspage",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *StatuspageError) Error() string {
	return e.Msg
}
