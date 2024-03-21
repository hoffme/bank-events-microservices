package infrastructure

import (
	"github.com/hoffme/backend-transactions/internal/domain"
	"github.com/hoffme/backend-transactions/internal/domain/transations/services/transactions"

	"github.com/hoffme/backend-transactions/internal/infrastructure/mongodb"
)

func (i Infrastructure) DomainDependencies() domain.Dependencies {
	accountRepository := mongodb.NewAccountRepository(i.mongo, i.rabbit)
	transactionRepository := mongodb.NewTransactionRepository(i.mongo, i.rabbit)
	transactionViewRepository := mongodb.NewTransactionViewRepository(i.mongo)

	return domain.Dependencies{
		Repositories: domain.Repositories{
			Account:         accountRepository,
			Transaction:     transactionRepository,
			TransactionView: transactionViewRepository,
		},
		Services: domain.Services{
			Transactions: transactions.NewService(accountRepository, transactionRepository),
		},
	}
}
