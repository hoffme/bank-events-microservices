package handlers

import (
	"context"

	"github.com/hoffme/backend-transactions/internal/domain"
	"github.com/hoffme/backend-transactions/internal/shared/events"
	"github.com/hoffme/backend-transactions/internal/shared/vo"

	accountsAccount "github.com/hoffme/backend-transactions/internal/domain/accounts/models/account"
)

var _ events.Handler = AccountUpdateName{}

type AccountUpdateName struct {
	header  events.HandlerHeader
	modules domain.Dependencies
}

func NewAccountUpdateName(modules domain.Dependencies) events.Handler {
	return AccountUpdateName{
		modules: modules,
		header: events.HandlerHeader{
			ID:            vo.UUIDRandom().Raw(),
			QueueTopic:    events.QueueTopic("account.update.name"),
			ConsumeTopics: []string{accountsAccount.EventNameChangedTopic},
			NoACK:         false,
			MaxIntents:    3,
		},
	}
}

func (s AccountUpdateName) Header() events.HandlerHeader {
	return s.header
}

func (s AccountUpdateName) Resolve(ctx context.Context, raw events.MessageRaw) error {
	evt, err := events.MessageFromRaw[accountsAccount.EventNameChangedParams](raw)
	if err != nil {
		return err
	}

	instance, err := s.modules.Repositories.Account.Get(ctx, evt.Payload.ID)
	if err != nil {
		return err
	}

	err = instance.SetName(evt.Payload.Name)
	if err != nil {
		return err
	}

	return s.modules.Repositories.Account.Save(ctx, instance)
}
