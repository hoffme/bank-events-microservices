package handlers

import (
	"context"

	"github.com/hoffme/backend-transactions/internal/domain"
	"github.com/hoffme/backend-transactions/internal/shared/events"
	"github.com/hoffme/backend-transactions/internal/shared/logger"
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

	// get models

	instance, err := s.modules.Repositories.Transaction.Get(ctx, evt.Payload.ID)
	if err != nil {
		return err
	}

	accountFrom, err := s.modules.Repositories.Account.Get(ctx, instance.GetFromAccountID())
	if err != nil {
		return err
	}

	accountTo, err := s.modules.Repositories.Account.Get(ctx, instance.GetToAccountID())
	if err != nil {
		return err
	}

	// make changes

	if !accountFrom.HasFounds(instance.GetAmount()) {
		err = instance.Reject()
		if err != nil {
			return err
		}

		return s.modules.Repositories.Transaction.Save(ctx, instance)
	}

	err = accountFrom.SetBalance(accountFrom.GetBalance() - instance.GetAmount())
	if err != nil {
		return err
	}

	err = accountTo.SetBalance(accountTo.GetBalance() + instance.GetAmount())
	if err != nil {
		return err
	}

	err = instance.Complete()
	if err != nil {
		return err
	}

	// apply changes

	err = s.modules.Repositories.Account.Save(ctx, accountFrom)
	if err != nil {
		logger.Error("saving account from authorize transaction: %s", err.Error())
		return err
	}

	err = s.modules.Repositories.Transaction.Save(ctx, instance)
	if err != nil {
		logger.Error("saving transaction: %s", err.Error())
		return err
	}

	err = s.modules.Repositories.Account.Save(ctx, accountTo)
	if err != nil {
		logger.Error("saving account to authorize transaction: %s", err.Error())
		return err
	}

	return nil
}
