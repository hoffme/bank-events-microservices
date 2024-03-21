package transactions

import (
	"github.com/hoffme/backend-transactions/internal/domain/transations/models/account"
	"github.com/hoffme/backend-transactions/internal/domain/transations/models/transaction"
)

type Service struct {
	accountsRepository     account.Repository
	transactionsRepository transaction.Repository
}

func NewService(account account.Repository, transactions transaction.Repository) Service {
	return Service{
		accountsRepository:     account,
		transactionsRepository: transactions,
	}
}
