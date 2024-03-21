package account

import "github.com/hoffme/backend-transactions/internal/shared/errors"

var Error = errors.New("TRANSACTIONS:ACCOUNT")

var ErrorNotFound = errors.New("TRANSACTIONS:ACCOUNT:NOT_FOUND").WithStatus(404)

var ErrorInsufficientFounds = errors.New("TRANSACTIONS:ACCOUNT:INSUFFICIENT_FOUNDS").WithStatus(404)
