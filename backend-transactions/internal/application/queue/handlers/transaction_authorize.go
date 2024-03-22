package handlers

import (
	"context"

	"github.com/hoffme/backend-transactions/internal/domain"
	"github.com/hoffme/backend-transactions/internal/shared/events"
	"github.com/hoffme/backend-transactions/internal/shared/vo"

	"github.com/hoffme/backend-transactions/internal/domain/transations/models/transaction"
)

var _ events.Handler = TransactionAuthorize{}

type TransactionAuthorize struct {
	header  events.HandlerHeader
	modules domain.Dependencies
}

func NewTransactionAuthorize(modules domain.Dependencies) events.Handler {
	return TransactionAuthorize{
		modules: modules,
		header: events.HandlerHeader{
			ID:            vo.UUIDRandom().Raw(),
			QueueTopic:    events.QueueTopic("transaction.authorize"),
			ConsumeTopics: []string{transaction.EventCreatedTopic},
			NoACK:         false,
			MaxIntents:    3,
		},
	}
}

func (s TransactionAuthorize) Header() events.HandlerHeader {
	return s.header
}

func (s TransactionAuthorize) Resolve(ctx context.Context, raw events.MessageRaw) error {
	evt, err := events.MessageFromRaw[transaction.EventCreatedParams](raw)
	if err != nil {
		return err
	}

	return s.modules.Services.Transactions.Authorize(ctx, evt.Payload)
}
