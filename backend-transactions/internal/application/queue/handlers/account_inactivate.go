package handlers

import (
	"context"

	"github.com/hoffme/backend-transactions/internal/domain"
	"github.com/hoffme/backend-transactions/internal/shared/events"
	"github.com/hoffme/backend-transactions/internal/shared/vo"

	accountsAccount "github.com/hoffme/backend-transactions/internal/domain/accounts/models/account"
)

var _ events.Handler = AccountInactivate{}

type AccountInactivate struct {
	header  events.HandlerHeader
	modules domain.Dependencies
}

func NewAccountInactivate(modules domain.Dependencies) events.Handler {
	return AccountInactivate{
		modules: modules,
		header: events.HandlerHeader{
			ID:            vo.UUIDRandom().Raw(),
			QueueTopic:    events.QueueTopic("account.inactivate"),
			ConsumeTopics: []string{accountsAccount.EventInactivatedTopic},
			NoACK:         false,
			MaxIntents:    3,
		},
	}
}

func (s AccountInactivate) Header() events.HandlerHeader {
	return s.header
}

func (s AccountInactivate) Resolve(ctx context.Context, raw events.MessageRaw) error {
	evt, err := events.MessageFromRaw[accountsAccount.EventInactivatedParams](raw)
	if err != nil {
		return err
	}

	instance, err := s.modules.Repositories.Account.Get(ctx, evt.Payload.ID)
	if err != nil {
		return err
	}

	err = instance.Deactivate()
	if err != nil {
		return err
	}

	return s.modules.Repositories.Account.Save(ctx, instance)
}
