package transaction

import "github.com/hoffme/backend-transactions/internal/shared/errors"

var Error = errors.New("TRANSACTIONS:TRANSACTION")

var ErrorNotFound = errors.New("TRANSACTIONS:TRANSACTION:NOT_FOUND").WithStatus(404)

var ErrorTransactionStateNotPending = errors.New("TRANSACTIONS:TRANSACTION:STATE:NOT_PENDING").WithStatus(404)
