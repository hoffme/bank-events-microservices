package events

import (
	"context"
)

// Bus

type Bus interface {
	Dispatcher
	Subscriber
}

type Dispatcher interface {
	Dispatch(context.Context, ...MessageRaw)
}

type Subscriber interface {
	Subscribe(context.Context, ...Handler) error
	Unsubscribe(context.Context, ...Handler) error
}

type HandlerHeader struct {
	ID            string
	QueueTopic    string
	ConsumeTopics []string
	NoACK         bool
	MaxIntents    int64
	TTL           int64
}

type Handler interface {
	Header() HandlerHeader
	Resolve(context.Context, MessageRaw) error
}
