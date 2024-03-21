package vo

import "github.com/hoffme/backend-transactions/internal/shared/errors"

// implementation

type Number struct {
	raw int64
}

func (n Number) Raw() int64 { return n.raw }

func (n Number) Equal(other Number) bool {
	return n.raw == other.Raw()
}

func (n Number) Lower(other Number) bool {
	return n.raw < other.Raw()
}

func (n Number) Grater(other Number) bool {
	return n.raw > other.Raw()
}

func (n Number) Add(other Number) Number {
	return Number{raw: n.raw + other.raw}
}

func (n Number) Sub(other Number) Number {
	return Number{raw: n.raw - other.raw}
}

// builders

var ErrorNumberInvalid = errors.New("VO:NUMBER:INVALID").WithStatus(400)

func NumberFromRaw(value int64) (Number, error) {
	return Number{raw: value}, nil
}

func NumberZero() Number {
	return Number{raw: 0}
}
