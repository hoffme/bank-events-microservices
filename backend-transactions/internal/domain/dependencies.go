package domain

import (
	"github.com/hoffme/backend-transactions/internal/domain/transations/models/account"
	"github.com/hoffme/backend-transactions/internal/domain/transations/models/transaction"
	"github.com/hoffme/backend-transactions/internal/domain/transations/models/transaction_view"
	"github.com/hoffme/backend-transactions/internal/domain/transations/services/transactions"
)

type Dependencies struct {
	Repositories Repositories
	Services     Services
}

type Repositories struct {
	Account         account.Repository
	Transaction     transaction.Repository
	TransactionView transaction_view.Repository
}

type Services struct {
	Transactions transactions.Service
}
