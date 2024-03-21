package handlers

import (
	"context"

	"github.com/hoffme/backend-transactions/internal/domain"
	"github.com/hoffme/backend-transactions/internal/shared/events"
	"github.com/hoffme/backend-transactions/internal/shared/vo"

	accountsAccount "github.com/hoffme/backend-transactions/internal/domain/accounts/models/account"
	transactionsAccount "github.com/hoffme/backend-transactions/internal/domain/transations/models/account"
)

var _ events.Handler = AccountCreate{}

type AccountCreate struct {
	header  events.HandlerHeader
	modules domain.Dependencies
}

func NewAccountCreate(modules domain.Dependencies) events.Handler {
	return AccountCreate{
		modules: modules,
		header: events.HandlerHeader{
			ID:            vo.UUIDRandom().Raw(),
			QueueTopic:    events.QueueTopic("account.create"),
			ConsumeTopics: []string{accountsAccount.EventCreatedTopic},
			NoACK:         false,
			MaxIntents:    3,
		},
	}
}

func (s AccountCreate) Header() events.HandlerHeader {
	return s.header
}

func (s AccountCreate) Resolve(ctx context.Context, raw events.MessageRaw) error {
	evt, err := events.MessageFromRaw[accountsAccount.EventCreatedParams](raw)
	if err != nil {
		return err
	}

	instance, err := transactionsAccount.Create(
		evt.Payload.ID,
		evt.Payload.Name,
		evt.Payload.Currency,
		evt.Payload.Balance,
		evt.Payload.Active,
	)
	if err != nil {
		return err
	}

	return s.modules.Repositories.Account.Save(ctx, instance)
}
