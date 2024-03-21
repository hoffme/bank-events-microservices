package vo

import "github.com/hoffme/backend-transactions/internal/shared/errors"

// implementation

type Boolean struct {
	raw bool
}

func (t Boolean) IsFalse() bool { return t.raw == false }

func (t Boolean) IsTrue() bool { return t.raw == true }

func (t Boolean) Equal(other Boolean) bool { return t.Raw() == other.Raw() }

func (t Boolean) Raw() bool { return t.raw }

// errors

var ErrorBoolInvalid = errors.New("VO:BOOL:INVALID").WithStatus(400)

// builders

func BoolFromRaw(value bool) (Boolean, error) {
	return Boolean{raw: value}, nil
}

func BoolFromRawString(value string) (Boolean, error) {
	if value != "true" && value != "false" {
		return Boolean{}, ErrorBoolInvalid
	}

	return Boolean{raw: value == "true"}, nil
}

func BoolFromRawInt(value int) (Boolean, error) {
	return Boolean{raw: value != 0}, nil
}

// constants

var BoolTrue = Boolean{raw: true}

var BoolFalse = Boolean{raw: false}
