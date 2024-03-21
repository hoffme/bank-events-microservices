package errors

import (
	"errors"

	"github.com/hoffme/backend-transactions/internal/shared/null"
)

var _ error = Error{}

type Error struct {
	status      int
	code        string
	description string
	wrap        null.Null[error]
}

func New(code string) Error {
	return Error{code: code}
}

func Wrap(err error) Error {
	return Error{
		wrap: null.WithValue(err),
	}
}

func (e Error) WithCode(code string) Error {
	e.code = code
	return e
}

func (e Error) WithStatus(status int) Error {
	e.status = status
	return e
}

func (e Error) WithDescription(description string) Error {
	e.description = description
	return e
}

func (e Error) Code() string {
	code := e.code

	if e.wrap.IsNotNull() {
		var appErr Error
		if errors.As(e.wrap.Value(), &appErr) {
			if len(code) > 0 {
				code += ":"
			}
			code += appErr.Code()
		}
	}

	return code
}

func (e Error) Status() int {
	if e.status > 0 {
		return e.status
	}

	if e.wrap.IsNotNull() {
		var appErr Error
		if errors.As(e.wrap.Value(), &appErr) {
			return appErr.Status()
		}
	}

	return 0
}

func (e Error) Description() string {
	if len(e.description) > 0 {
		return e.description
	}

	if e.wrap.IsNotNull() {
		var appErr Error
		if errors.As(e.wrap.Value(), &appErr) {
			return appErr.Description()
		}
	}

	return ""
}

func (e Error) Error() string { return e.Code() }
