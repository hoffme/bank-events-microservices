package transactions

import (
	"context"

	"github.com/hoffme/backend-transactions/internal/domain/transations/models/transaction"
	"github.com/hoffme/backend-transactions/internal/shared/logger"
)

func (s Service) Authorize(ctx context.Context, evt transaction.EventCreatedParams) error {
	// get models

	instance, err := s.transactionsRepository.Get(ctx, evt.ID)
	if err != nil {
		return err
	}

	accountFrom, err := s.accountsRepository.Get(ctx, instance.GetFromAccountID())
	if err != nil {
		return err
	}

	accountTo, err := s.accountsRepository.Get(ctx, instance.GetToAccountID())
	if err != nil {
		return err
	}

	// make changes

	if !accountFrom.HasFounds(instance.GetAmount()) {
		err = instance.Reject()
		if err != nil {
			return err
		}

		return s.transactionsRepository.Save(ctx, instance)
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
	// TODO: save this in transactional

	err = s.accountsRepository.Save(ctx, accountFrom)
	if err != nil {
		logger.Error("saving account from authorize transaction: %s", err.Error())
		return err
	}

	err = s.transactionsRepository.Save(ctx, instance)
	if err != nil {
		logger.Error("saving transaction: %s", err.Error())
		return err
	}

	err = s.accountsRepository.Save(ctx, accountTo)
	if err != nil {
		logger.Error("saving account to authorize transaction: %s", err.Error())
		return err
	}

	return nil
}
