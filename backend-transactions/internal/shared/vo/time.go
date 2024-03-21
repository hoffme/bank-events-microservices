package vo

import (
	"time"

	"github.com/hoffme/backend-transactions/internal/shared/errors"
)

// implementation

type Time struct {
	raw time.Time
}

func (t Time) Equal(t2 Time) bool {
	return t.raw.Equal(t2.Raw())
}

func (t Time) Before(t2 Time) bool {
	return t.raw.Before(t2.Raw())
}

func (t Time) After(t2 Time) bool {
	return t.raw.After(t2.Raw())
}

func (t Time) Raw() time.Time {
	return t.raw
}

// builders

var ErrorTimeInvalid = errors.New("VO:TIME:INVALID").WithStatus(400)

func TimeNow() Time {
	return Time{raw: time.Now()}
}

func TimeFromRaw(value time.Time) (Time, error) {
	if value.IsZero() {
		return Time{}, ErrorTimeInvalid
	}

	return Time{raw: value}, nil
}
