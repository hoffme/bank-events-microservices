package transactions

import (
	"github.com/hoffme/backend-transactions/internal/shared/errors"
)

var ErrorAmountNotPositive = errors.New("TRANSACTIONS:TRANSFER:AMOUNT:NOT_POSITIVE")

var ErrorTransferCreated = errors.New("TRANSACTIONS:TRANSFER:CREATED")

var ErrorTransferAccountFrom = errors.New("TRANSACTIONS:TRANSFER:ACCOUNT_FROM")

var ErrorTransferAccountTo = errors.New("TRANSACTIONS:TRANSFER:ACCOUNT_TO")
