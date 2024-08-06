package appError

import (
	"fmt"
	"runtime"

	errorCode "banana-account-book.com/internal/libs/app-error/error-code"
)

type applicationError struct {
	Message       string
	Code          int
	ClientMessage string
	Stack         string
}

func New(msg string, code int, clientMsg string) *applicationError {
	err := applicationError{
		Message:       msg,
		Code:          code,
		Stack:         fmt.Sprintf("Error: %s", msg),
		ClientMessage: clientMsg,
	}
	return err.stackTrace()
}

func (e applicationError) Error() string {
	return e.Message
}

var funcInfoFormat = "Stack Trace: {%s:%d} [%s]"

func getFuncInfo(pc uintptr, file string, line int) string {
	f := runtime.FuncForPC(pc)
	if f == nil {
		return fmt.Sprintf(funcInfoFormat, file, line, "unknwon")
	}
	return fmt.Sprintf(funcInfoFormat, file, line, f.Name())
}

var wrapFormat = "%s\n%s" // "error \n {file:line} [func name] msg"

func (e *applicationError) stackTrace() *applicationError {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		e.Stack = fmt.Sprintf(wrapFormat, e.Stack, e.Error())
	}
	e.Stack = fmt.Sprintf(wrapFormat, e.Stack, getFuncInfo(pc, file, line))
	return e
}

func Wrap(err error) error {
	if e, ok := err.(*applicationError); ok {
		return e.stackTrace()
	}
	// NOTE: Set status with 500 when error is not application error
	return New(err.Error(), errorCode.Internal, "Internal Server Error").stackTrace()
}

func UnWrap(err error) *applicationError {
	if e, ok := err.(*applicationError); ok {
		return e
	}
	// NOTE: Set status with 500 when error is not application error
	return New(err.Error(), errorCode.Internal, "Internal Server Error").stackTrace()
}
