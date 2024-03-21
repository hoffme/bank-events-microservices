package vo

import (
	"github.com/hoffme/backend-transactions/internal/shared/errors"

	"github.com/google/uuid"
)

// implementation

type UUID struct {
	raw string
}

func (i UUID) Raw() string {
	return i.raw
}

func (i UUID) Equal(other UUID) bool {
	return i.raw == other.Raw()
}

// builders

var ErrorUUIDInvalid = errors.New("VO:UUID:INVALID").WithStatus(400)

func UUIDFromRaw(value string) (UUID, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return UUID{}, ErrorUUIDInvalid
	}

	return UUID{raw: id.String()}, nil
}

func UUIDRandom() UUID {
	return UUID{raw: uuid.NewString()}
}
