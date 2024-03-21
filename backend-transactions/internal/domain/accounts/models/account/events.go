package account

import "github.com/hoffme/backend-transactions/internal/shared/events"

// created

var EventCreatedTopic = "app.accounts.evt.account.created"

var _ events.Payload = EventCreatedParams{}

type EventCreatedParams struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Active    bool   `json:"active"`
	Balance   int64  `json:"balance"`
	Currency  string `json:"currency"`
	CreatedAt string `json:"created_at"`
}

func (e EventCreatedParams) Topic() string { return EventCreatedTopic }

// name changed

var EventNameChangedTopic = "app.accounts.evt.account.name.changed"

var _ events.Payload = EventNameChangedParams{}

type EventNameChangedParams struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (e EventNameChangedParams) Topic() string { return EventNameChangedTopic }

// activated

var EventActivatedTopic = "app.accounts.evt.account.activated"

var _ events.Payload = EventActivatedParams{}

type EventActivatedParams struct {
	ID string `json:"id"`
}

func (e EventActivatedParams) Topic() string { return EventActivatedTopic }

// inactivate

var EventInactivatedTopic = "app.accounts.evt.account.inactivate"

var _ events.Payload = EventInactivatedParams{}

type EventInactivatedParams struct {
	ID string `json:"id"`
}

func (e EventInactivatedParams) Topic() string { return EventInactivatedTopic }
