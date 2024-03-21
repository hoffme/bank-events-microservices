package mongodb

import (
	"context"

	"github.com/hoffme/backend-transactions/internal/shared/events"

	"github.com/hoffme/backend-transactions/internal/domain/transations/models/account"
)

var _ account.Repository = AccountRepository{}

type AccountRepository struct {
	repo       repo[recordAccount]
	dispatcher events.Dispatcher
}

func NewAccountRepository(connection *Connection, dispatcher events.Dispatcher) AccountRepository {
	return AccountRepository{
		repo: repo[recordAccount]{
			connection:     connection,
			databaseName:   "bank_transactions",
			collectionName: "accounts",
			notFoundError:  account.ErrorNotFound,
			fieldsProperties: map[string]string{
				"id":       "_id",
				"name":     "name",
				"balance":  "balance",
				"currency": "currency",
				"active":   "active",
			},
		},
		dispatcher: dispatcher,
	}
}

func (a AccountRepository) Get(ctx context.Context, id string) (*account.Aggregate, error) {
	result, err := a.repo.get(ctx, id)
	if err != nil {
		return nil, err
	}

	return result.toRaw().Instance()
}

func (a AccountRepository) Save(ctx context.Context, instance *account.Aggregate) error {
	record := recordAccount{}.fromRaw(instance.Raw())

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

type recordAccount struct {
	ID       string `bson:"_id"`
	Name     string `bson:"name"`
	Balance  int64  `bson:"balance"`
	Currency string `bson:"currency"`
	Active   bool   `bson:"active"`
}

func (r recordAccount) fromRaw(raw account.Raw) recordAccount {
	return recordAccount{
		ID:       raw.Root.ID,
		Name:     raw.Root.Name,
		Balance:  raw.Root.Balance,
		Currency: raw.Root.Currency,
		Active:   raw.Root.Active,
	}
}

func (r recordAccount) toRaw() account.Raw {
	return account.Raw{
		Root: account.RootRaw{
			ID:       r.ID,
			Name:     r.Name,
			Balance:  r.Balance,
			Currency: r.Currency,
			Active:   r.Active,
		},
	}
}
