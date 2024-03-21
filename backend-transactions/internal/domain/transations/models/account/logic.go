package account

import (
	"github.com/hoffme/backend-transactions/internal/shared/errors"
	"github.com/hoffme/backend-transactions/internal/shared/events"
	"github.com/hoffme/backend-transactions/internal/shared/vo"
)

func Create(id, name, currency string, balance int64, active bool) (*Aggregate, error) {
	raw := Raw{
		Root: RootRaw{
			ID:       id,
			Name:     name,
			Balance:  balance,
			Currency: currency,
			Active:   active,
		},
	}

	instance, err := raw.Instance()
	if err != nil {
		return nil, err
	}

	instance.events.Record(events.NewMessageGeneric(EventCreatedParams{
		ID:       instance.GetID(),
		Name:     instance.GetName(),
		Balance:  instance.GetBalance(),
		Currency: instance.GetCurrency(),
		Active:   instance.GetActive(),
	}))

	return instance, nil
}

// getters

func (a *Aggregate) GetID() string {
	return a.root.ID.Raw()
}

func (a *Aggregate) GetName() string {
	return a.root.Name.Raw()
}

func (a *Aggregate) GetBalance() int64 {
	return a.root.Balance.Raw()
}

func (a *Aggregate) GetCurrency() string {
	return a.root.Currency.Raw()
}

func (a *Aggregate) GetActive() bool {
	return a.root.Active.Raw()
}

func (a *Aggregate) IsActive() bool {
	return a.root.Active.IsTrue()
}

func (a *Aggregate) IsInactive() bool {
	return a.root.Active.IsFalse()
}

// setters

func (a *Aggregate) SetName(nameRaw string) error {
	name, err := vo.StringFromRaw(nameRaw)
	if err != nil {
		return errors.Wrap(err).WithCode(Error.Code())
	}

	if a.root.Name.Equal(name) {
		return nil
	}

	a.root.Name = name

	a.events.Record(events.NewMessageGeneric(EventNameChangedParams{
		ID:   a.GetID(),
		Name: a.GetName(),
	}))

	return nil
}

func (a *Aggregate) SetBalance(balanceRaw int64) error {
	balance, err := vo.NumberFromRaw(balanceRaw)
	if err != nil {
		return errors.Wrap(err).WithCode(Error.Code())
	}

	if balance.Lower(vo.NumberZero()) {
		return ErrorInsufficientFounds
	}

	a.root.Balance = balance

	a.events.Record(events.NewMessageGeneric(EventBalanceChangedParams{
		ID:      a.GetID(),
		Balance: a.GetBalance(),
	}))

	return nil
}

func (a *Aggregate) IncrementBalance(factor int64) error {
	return a.SetBalance(a.GetBalance() + factor)
}

func (a *Aggregate) DecrementBalance(factor int64) error {
	return a.SetBalance(a.GetBalance() - factor)
}

func (a *Aggregate) Deactivate() error {
	if a.root.Active.IsFalse() {
		return nil
	}

	a.root.Active = vo.BoolFalse

	a.events.Record(events.NewMessageGeneric(EventInactivatedParams{
		ID: a.GetID(),
	}))

	return nil
}

func (a *Aggregate) Activate() error {
	if a.root.Active.IsTrue() {
		return nil
	}

	a.root.Active = vo.BoolTrue

	a.events.Record(events.NewMessageGeneric(EventActivatedParams{
		ID: a.GetID(),
	}))

	return nil
}

// logic

func (a *Aggregate) HasFounds(amountRaw int64) bool {
	amount, err := vo.NumberFromRaw(amountRaw)
	if err != nil {
		return false
	}

	return a.root.Balance.Grater(amount) || a.root.Balance.Equal(amount)
}

// events

func (a *Aggregate) PullEvents() []events.MessageRaw {
	return a.events.Pull()
}
