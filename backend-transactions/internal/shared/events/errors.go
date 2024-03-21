package events

import (
	"errors"
)

var ErrorInternal = errors.New("EVENTS:INTERNAL")

var ErrorTopicInvalid = errors.New("EVENTS:TOPIC:INVALID")
