package handlers

import (
	"context"

	"github.com/hoffme/backend-transactions/internal/domain"
	"github.com/hoffme/backend-transactions/internal/shared/events"
	"github.com/hoffme/backend-transactions/internal/shared/vo"

	accountsAccount "github.com/hoffme/backend-transactions/internal/domain/accounts/models/account"
)

var _ events.Handler = AccountActivate{}

type AccountActivate struct {
	header  events.HandlerHeader
	modules domain.Dependencies
}

func NewAccountActivate(modules domain.Dependencies) events.Handler {
	return AccountActivate{
		modules: modules,
		header: events.HandlerHeader{
			ID:            vo.UUIDRandom().Raw(),
			QueueTopic:    events.QueueTopic("account.activate"),
			ConsumeTopics: []string{accountsAccount.EventActivatedTopic},
			NoACK:         false,
			MaxIntents:    3,
		},
	}
}

func (s AccountActivate) Header() events.HandlerHeader {
	return s.header
}

func (s AccountActivate) Resolve(ctx context.Context, raw events.MessageRaw) error {
	evt, err := events.MessageFromRaw[accountsAccount.EventActivatedParams](raw)
	if err != nil {
		return err
	}

	instance, err := s.modules.Repositories.Account.Get(ctx, evt.Payload.ID)
	if err != nil {
		return err
	}

	err = instance.Activate()
	if err != nil {
		return err
	}

	return s.modules.Repositories.Account.Save(ctx, instance)
}
