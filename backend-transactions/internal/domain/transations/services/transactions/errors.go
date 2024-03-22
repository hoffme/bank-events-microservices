package transactions

import (
	"github.com/hoffme/backend-transactions/internal/shared/errors"
)

var ErrorAmountNotPositive = errors.New("TRANSACTIONS:TRANSFER:AMOUNT:NOT_POSITIVE")

var ErrorTransferCreated = errors.New("TRANSACTIONS:TRANSFER:CREATED")

var ErrorTransferAccountFromInsufficientsFound = errors.New("TRANSACTIONS:TRANSFER:ACCOUNT_FROM:INSSUFICIENTS_FOUNDS")

var ErrorTransferAccountFromInactive = errors.New("TRANSACTIONS:TRANSFER:ACCOUNT_FROM:INACTIVE")

var ErrorTransferAccountToInactive = errors.New("TRANSACTIONS:TRANSFER:ACCOUNT_TO:INACTIVE")
