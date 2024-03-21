package transactions

import (
	"github.com/hoffme/backend-transactions/internal/shared/vo"
)

type TransferParamsRaw struct {
	TransferID    string
	FromAccountID string
	ToAccountID   string
	Currency      string
	Amount        int64
}

func (raw TransferParamsRaw) toEntity() (TransferParamsEntity, error) {
	var err error

	result := TransferParamsEntity{}

	result.TransferID, err = vo.UUIDFromRaw(raw.TransferID)
	if err != nil {
		return TransferParamsEntity{}, err
	}

	result.FromAccountID, err = vo.UUIDFromRaw(raw.FromAccountID)
	if err != nil {
		return TransferParamsEntity{}, err
	}

	result.ToAccountID, err = vo.UUIDFromRaw(raw.ToAccountID)
	if err != nil {
		return TransferParamsEntity{}, err
	}

	result.Currency, err = vo.CurrencyFromRaw(raw.Currency)
	if err != nil {
		return TransferParamsEntity{}, err
	}

	result.Amount, err = vo.NumberFromRaw(raw.Amount)
	if err != nil {
		return TransferParamsEntity{}, err
	}
	if result.Amount.Lower(vo.NumberZero()) || result.Amount.Equal(vo.NumberZero()) {
		return TransferParamsEntity{}, ErrorAmountNotPositive
	}

	return result, nil
}

type TransferParamsEntity struct {
	TransferID    vo.UUID
	FromAccountID vo.UUID
	ToAccountID   vo.UUID
	Currency      vo.Currency
	Amount        vo.Number
}
