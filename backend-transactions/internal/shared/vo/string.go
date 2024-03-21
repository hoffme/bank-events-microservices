package vo

import "github.com/hoffme/backend-transactions/internal/shared/errors"

// implementation

type String struct {
	raw string
}

func (t String) Equal(other String) bool { return t.Raw() == other.Raw() }

func (t String) Length() int { return len(t.raw) }

func (t String) IsEmpty() bool { return t.Length() == 0 }

func (t String) IsNotEmpty() bool { return t.Length() > 0 }

func (t String) Raw() string { return t.raw }

// builders

var ErrorStringInvalid = errors.New("VO:STRING:INVALID").WithStatus(400)

func StringFromRaw(value string) (String, error) {
	return String{raw: value}, nil
}
