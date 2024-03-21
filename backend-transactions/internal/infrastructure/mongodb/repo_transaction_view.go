package mongodb

import (
	"context"
	"time"

	"github.com/hoffme/backend-transactions/internal/domain/transations/models/transaction_view"
	"github.com/hoffme/backend-transactions/internal/shared/repository"

	"github.com/hoffme/backend-transactions/internal/shared/null"

	"github.com/hoffme/backend-transactions/internal/domain/transations/models/transaction"
)

var _ transaction_view.Repository = TransactionViewRepository{}

type TransactionViewRepository struct {
	repo repo[recordTransactionView]
}

func NewTransactionViewRepository(connection *Connection) TransactionViewRepository {
	return TransactionViewRepository{
		repo: repo[recordTransactionView]{
			connection:     connection,
			databaseName:   "bank_transactions",
			collectionName: "view_transactions",
			notFoundError:  transaction.ErrorNotFound,
			fieldsProperties: map[string]string{
				"id":                "_id",
				"from_account.id":   "from_account_id",
				"from_account.name": "from_account_name",
				"to_account.id":     "to_account_id",
				"to_account.name":   "to_account_name",
				"state":             "state",
				"amount":            "amount",
				"currency":          "currency",
				"created_at":        "created_at",
				"finished_at":       "finished_at",
			},
		},
	}
}

func (a TransactionViewRepository) Get(ctx context.Context, id string) (*transaction_view.Aggregate, error) {
	result, err := a.repo.get(ctx, id)
	if err != nil {
		return nil, err
	}

	return result.toRaw().Instance()
}

func (a TransactionViewRepository) Search(ctx context.Context, filter repository.SearchFilter) (repository.SearchResult[*transaction_view.Aggregate], error) {
	result, err := a.repo.search(ctx, filter)
	if err != nil {
		return repository.SearchResult[*transaction_view.Aggregate]{}, err
	}

	return repository.MapResultWithError(result, func(record recordTransactionView) (*transaction_view.Aggregate, error) {
		return record.toRaw().Instance()
	})
}

// utils

type recordTransactionView struct {
	ID              string               `bson:"_id"`
	FromAccountID   string               `bson:"from_account_id"`
	FromAccountName string               `bson:"from_account_name"`
	ToAccountID     string               `bson:"to_account_id"`
	ToAccountName   string               `bson:"to_account_name"`
	State           string               `bson:"state"`
	Amount          int64                `bson:"amount"`
	Currency        string               `bson:"currency"`
	CreatedAt       time.Time            `bson:"created_at"`
	FinishedAt      null.Null[time.Time] `bson:"finished_at"`
}

func (r recordTransactionView) fromRaw(raw transaction_view.Raw) recordTransactionView {
	return recordTransactionView{
		ID:              raw.Root.ID,
		FromAccountID:   raw.From.ID,
		FromAccountName: raw.From.Name,
		ToAccountID:     raw.To.ID,
		ToAccountName:   raw.To.Name,
		State:           raw.Root.State,
		Amount:          raw.Root.Amount,
		Currency:        raw.Root.Currency,
		CreatedAt:       raw.Root.CreatedAt,
		FinishedAt:      raw.Root.FinishedAt,
	}
}

func (r recordTransactionView) toRaw() transaction_view.Raw {
	return transaction_view.Raw{
		Root: transaction_view.RootRaw{
			ID:         r.ID,
			State:      r.State,
			Amount:     r.Amount,
			Currency:   r.Currency,
			CreatedAt:  r.CreatedAt,
			FinishedAt: r.FinishedAt,
		},
		From: transaction_view.AccountRaw{
			ID:   r.FromAccountID,
			Name: r.FromAccountName,
		},
		To: transaction_view.AccountRaw{
			ID:   r.ToAccountID,
			Name: r.ToAccountName,
		},
	}
}
