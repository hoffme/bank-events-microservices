package transaction

import (
	"time"

	"github.com/hoffme/backend-transactions/internal/shared/events"
	"github.com/hoffme/backend-transactions/internal/shared/null"
	"github.com/hoffme/backend-transactions/internal/shared/vo"
)

func Create(id, fromAccountId, toAccountId, currency string, amount int64) (*Aggregate, error) {
	raw := Raw{
		Root: RootRaw{
			ID:            id,
			FromAccountID: fromAccountId,
			ToAccountID:   toAccountId,
			State:         vo.TransactionStatePending().Raw(),
			Amount:        amount,
			Currency:      currency,
			CreatedAt:     time.Now(),
			FinishedAt:    null.WithoutValue[time.Time](),
		},
	}

	instance, err := raw.Instance()
	if err != nil {
		return nil, err
	}

	instance.events.Record(events.NewMessageGeneric(EventCreatedParams{
		ID:            instance.GetID(),
		FromAccountID: instance.GetFromAccountID(),
		ToAccountID:   instance.GetToAccountID(),
		State:         instance.GetState(),
		Amount:        instance.GetAmount(),
		Currency:      instance.GetCurrency(),
		CreatedAt:     instance.GetCreatedAt(),
	}))

	return instance, nil
}

// getters

func (a *Aggregate) GetID() string {
	return a.root.ID.Raw()
}

func (a *Aggregate) GetFromAccountID() string {
	return a.root.FromAccountID.Raw()
}

func (a *Aggregate) GetToAccountID() string {
	return a.root.ToAccountID.Raw()
}

func (a *Aggregate) GetState() string {
	return a.root.State.Raw()
}

func (a *Aggregate) GetCurrency() string {
	return a.root.Currency.Raw()
}

func (a *Aggregate) GetAmount() int64 {
	return a.root.Amount.Raw()
}

func (a *Aggregate) GetCreatedAt() time.Time {
	return a.root.CreatedAt.Raw()
}

func (a *Aggregate) GetFinishedAt() null.Null[time.Time] {
	return null.Raw(a.root.FinishedAt)
}

// setters

func (a *Aggregate) Complete() error {
	if !a.root.State.Equal(vo.TransactionStatePending()) {
		return ErrorTransactionStateNotPending
	}

	a.root.State = vo.TransactionStateCompleted()
	a.root.FinishedAt = null.WithValue(vo.TimeNow())

	a.events.Record(events.NewMessageGeneric(EventStateCompletedParams{
		ID:         a.GetID(),
		State:      a.GetState(),
		FinishedAt: a.GetFinishedAt().Value(),
	}))

	return nil
}

func (a *Aggregate) Reject() error {
	if !a.root.State.Equal(vo.TransactionStatePending()) {
		return ErrorTransactionStateNotPending
	}

	a.root.State = vo.TransactionStateRejected()
	a.root.FinishedAt = null.WithValue(vo.TimeNow())

	a.events.Record(events.NewMessageGeneric(EventStateRejectedParams{
		ID:         a.GetID(),
		State:      a.GetState(),
		FinishedAt: a.GetFinishedAt().Value(),
	}))

	return nil
}

// events

func (a *Aggregate) PullEvents() []events.MessageRaw {
	return a.events.Pull()
}
