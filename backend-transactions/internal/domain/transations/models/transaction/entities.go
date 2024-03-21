package transaction

import (
	"time"

	"github.com/hoffme/backend-transactions/internal/shared/errors"
	"github.com/hoffme/backend-transactions/internal/shared/null"
	"github.com/hoffme/backend-transactions/internal/shared/vo"
)

// Root

type RootRaw struct {
	ID            string
	FromAccountID string
	ToAccountID   string
	State         string
	Amount        int64
	Currency      string
	CreatedAt     time.Time
	FinishedAt    null.Null[time.Time]
}

func (r RootRaw) Entity() (RootEntity, error) {
	var err error

	entity := RootEntity{}

	entity.ID, err = vo.UUIDFromRaw(r.ID)
	if err != nil {
		return RootEntity{}, errors.Wrap(err).WithCode("ID")
	}

	entity.FromAccountID, err = vo.UUIDFromRaw(r.FromAccountID)
	if err != nil {
		return RootEntity{}, errors.Wrap(err).WithCode("FROM_ACCOUNT_ID")
	}

	entity.ToAccountID, err = vo.UUIDFromRaw(r.ToAccountID)
	if err != nil {
		return RootEntity{}, errors.Wrap(err).WithCode("TO_ACCOUNT_ID")
	}

	entity.State, err = vo.TransactionStateFromRaw(r.State)
	if err != nil {
		return RootEntity{}, errors.Wrap(err).WithCode("STATE")
	}

	entity.Amount, err = vo.NumberFromRaw(r.Amount)
	if err != nil {
		return RootEntity{}, errors.Wrap(err).WithCode("AMOUNT")
	}

	entity.Currency, err = vo.CurrencyFromRaw(r.Currency)
	if err != nil {
		return RootEntity{}, errors.Wrap(err).WithCode("CURRENCY")
	}

	entity.CreatedAt, err = vo.TimeFromRaw(r.CreatedAt)
	if err != nil {
		return RootEntity{}, errors.Wrap(err).WithCode("CREATED_AT")
	}

	entity.FinishedAt, err = null.WithError(r.FinishedAt, vo.TimeFromRaw)
	if err != nil {
		return RootEntity{}, errors.Wrap(err).WithCode("FINISHED_AT")
	}

	return entity, err
}

type RootEntity struct {
	ID            vo.UUID
	FromAccountID vo.UUID
	ToAccountID   vo.UUID
	State         vo.TransactionState
	Amount        vo.Number
	Currency      vo.Currency
	CreatedAt     vo.Time
	FinishedAt    null.Null[vo.Time]
}

func (e RootEntity) Raw() RootRaw {
	return RootRaw{
		ID:            e.ID.Raw(),
		FromAccountID: e.FromAccountID.Raw(),
		ToAccountID:   e.ToAccountID.Raw(),
		State:         e.State.Raw(),
		Amount:        e.Amount.Raw(),
		Currency:      e.Currency.Raw(),
		CreatedAt:     e.CreatedAt.Raw(),
		FinishedAt:    null.Raw(e.FinishedAt),
	}
}
