package account

import (
	"github.com/hoffme/backend-transactions/internal/shared/events"
)

// Created

var EventCreatedTopic = events.EventTopic("transactions.account.created")

var _ events.Payload = EventCreatedParams{}

type EventCreatedParams struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
	Active   bool   `json:"active"`
}

func (e EventCreatedParams) Topic() string { return EventCreatedTopic }

// NameChanged

var EventNameChangedTopic = events.EventTopic("transactions.account.name.changed")

var _ events.Payload = EventNameChangedParams{}

type EventNameChangedParams struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (e EventNameChangedParams) Topic() string { return EventNameChangedTopic }

// Activated

var EventActivatedTopic = events.EventTopic("transactions.account.activated")

var _ events.Payload = EventActivatedParams{}

type EventActivatedParams struct {
	ID string `json:"id"`
}

func (e EventActivatedParams) Topic() string { return EventActivatedTopic }

// Inactivated

var EventInactivatedTopic = events.EventTopic("transactions.account.inactivated")

var _ events.Payload = EventInactivatedParams{}

type EventInactivatedParams struct {
	ID string `json:"id"`
}

func (e EventInactivatedParams) Topic() string { return EventInactivatedTopic }

// BalanceChanged

var EventBalanceChangedTopic = events.EventTopic("transactions.account.balance.changed")

var _ events.Payload = EventBalanceChangedParams{}

type EventBalanceChangedParams struct {
	ID      string `json:"id"`
	Balance int64  `json:"balance"`
}

func (e EventBalanceChangedParams) Topic() string { return EventBalanceChangedTopic }
