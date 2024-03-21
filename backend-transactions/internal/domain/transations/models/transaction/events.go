package transaction

import (
	"time"

	"github.com/hoffme/backend-transactions/internal/shared/events"
)

// Created

var EventCreatedTopic = events.EventTopic("transactions.transaction.created")

var _ events.Payload = EventCreatedParams{}

type EventCreatedParams struct {
	ID            string    `json:"id"`
	FromAccountID string    `json:"from_account_id"`
	ToAccountID   string    `json:"to_account_id"`
	State         string    `json:"state"`
	Amount        int64     `json:"amount"`
	Currency      string    `json:"currency"`
	CreatedAt     time.Time `json:"created_at"`
}

func (e EventCreatedParams) Topic() string { return EventCreatedTopic }

// StateCompleted

var EventStateCompletedTopic = events.EventTopic("transactions.transaction.state.completed")

var _ events.Payload = EventStateCompletedParams{}

type EventStateCompletedParams struct {
	ID         string    `json:"id"`
	State      string    `json:"state"`
	FinishedAt time.Time `json:"finished_at"`
}

func (e EventStateCompletedParams) Topic() string { return EventStateCompletedTopic }

// StateRejected

var EventStateRejectedTopic = events.EventTopic("transactions.transaction.state.rejected")

var _ events.Payload = EventStateRejectedParams{}

type EventStateRejectedParams struct {
	ID         string    `json:"id"`
	State      string    `json:"state"`
	FinishedAt time.Time `json:"finished_at"`
}

func (e EventStateRejectedParams) Topic() string { return EventStateRejectedTopic }
