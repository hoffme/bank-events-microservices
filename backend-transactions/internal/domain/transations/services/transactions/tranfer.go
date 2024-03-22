package transactions

import (
	"context"
	"errors"

	"github.com/hoffme/backend-transactions/internal/domain/transations/models/transaction"
)

func (s Service) Transfer(ctx context.Context, paramsRaw TransferParamsRaw) error {
	params, err := paramsRaw.toEntity()
	if err != nil {
		return err
	}

	_, err = s.transactionsRepository.Get(ctx, params.TransferID.Raw())
	if err == nil {
		return ErrorTransferCreated
	}
	if !errors.Is(err, transaction.ErrorNotFound) {
		return err
	}

	accountFrom, err := s.accountsRepository.Get(ctx, params.FromAccountID.Raw())
	if err != nil {
		return err
	}
	if accountFrom.IsInactive() {
		return ErrorTransferAccountFromInactive
	}
	if !accountFrom.HasFounds(params.Amount.Raw()) {
		return ErrorTransferAccountFromInsufficientsFound
	}

	accountTo, err := s.accountsRepository.Get(ctx, params.ToAccountID.Raw())
	if err != nil {
		return err
	}
	if accountTo.IsInactive() {
		return ErrorTransferAccountToInactive
	}

	result, err := transaction.Create(
		params.TransferID.Raw(),
		accountFrom.GetID(),
		accountTo.GetID(),
		params.Currency.Raw(),
		params.Amount.Raw(),
	)
	if err != nil {
		return err
	}

	return s.transactionsRepository.Save(ctx, result)
}
