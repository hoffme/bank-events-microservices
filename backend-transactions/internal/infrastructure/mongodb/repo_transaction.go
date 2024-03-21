package mongodb

import (
	"context"
	"time"

	"github.com/hoffme/backend-transactions/internal/shared/events"
	"github.com/hoffme/backend-transactions/internal/shared/null"

	"github.com/hoffme/backend-transactions/internal/domain/transations/models/transaction"
)

var _ transaction.Repository = TransactionRepository{}

type TransactionRepository struct {
	repo       repo[recordTransaction]
	dispatcher events.Dispatcher
}

func NewTransactionRepository(connection *Connection, dispatcher events.Dispatcher) TransactionRepository {
	return TransactionRepository{
		repo: repo[recordTransaction]{
			connection:     connection,
			databaseName:   "bank_transactions",
			collectionName: "transactions",
			notFoundError:  transaction.ErrorNotFound,
			fieldsProperties: map[string]string{
				"id":              "_id",
				"from_account_id": "from_account_id",
				"to_account_id":   "to_account_id",
				"state":           "state",
				"amount":          "amount",
				"currency":        "currency",
				"created_at":      "created_at",
				"finished_at":     "finished_at",
			},
		},
		dispatcher: dispatcher,
	}
}

func (a TransactionRepository) Get(ctx context.Context, id string) (*transaction.Aggregate, error) {
	result, err := a.repo.get(ctx, id)
	if err != nil {
		return nil, err
	}

	return result.toRaw().Instance()
}

func (a TransactionRepository) Save(ctx context.Context, instance *transaction.Aggregate) error {
	record := recordTransaction{}.fromRaw(instance.Raw())

	_, err := a.repo.save(ctx, instance.GetID(), record)
	if err != nil {
		return err
	}

	evt := instance.PullEvents()
	if a.dispatcher != nil && len(evt) > 0 {
		a.dispatcher.Dispatch(ctx, evt...)
	}

	return nil
}

// utils

type recordTransaction struct {
	ID            string               `bson:"_id"`
	FromAccountID string               `bson:"from_account_id"`
	ToAccountID   string               `bson:"to_account_id"`
	State         string               `bson:"state"`
	Amount        int64                `bson:"amount"`
	Currency      string               `bson:"currency"`
	CreatedAt     time.Time            `bson:"created_at"`
	FinishedAt    null.Null[time.Time] `bson:"finished_at"`
}

func (r recordTransaction) fromRaw(raw transaction.Raw) recordTransaction {
	return recordTransaction{
		ID:            raw.Root.ID,
		FromAccountID: raw.Root.FromAccountID,
		ToAccountID:   raw.Root.ToAccountID,
		State:         raw.Root.State,
		Amount:        raw.Root.Amount,
		Currency:      raw.Root.Currency,
		CreatedAt:     raw.Root.CreatedAt,
		FinishedAt:    raw.Root.FinishedAt,
	}
}

func (r recordTransaction) toRaw() transaction.Raw {
	return transaction.Raw{
		Root: transaction.RootRaw{
			ID:            r.ID,
			FromAccountID: r.FromAccountID,
			ToAccountID:   r.ToAccountID,
			State:         r.State,
			Amount:        r.Amount,
			Currency:      r.Currency,
			CreatedAt:     r.CreatedAt,
			FinishedAt:    r.FinishedAt,
		},
	}
}
