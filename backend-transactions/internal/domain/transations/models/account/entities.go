package account

import (
	"github.com/hoffme/backend-transactions/internal/shared/errors"

	"github.com/hoffme/backend-transactions/internal/shared/vo"
)

// Root

type RootRaw struct {
	ID       string
	Name     string
	Balance  int64
	Currency string
	Active   bool
}

func (r RootRaw) Entity() (RootEntity, error) {
	var err error

	entity := RootEntity{}

	entity.ID, err = vo.UUIDFromRaw(r.ID)
	if err != nil {
		return RootEntity{}, errors.Wrap(err).WithCode("ID")
	}

	entity.Name, err = vo.StringFromRaw(r.Name)
	if err != nil {
		return RootEntity{}, errors.Wrap(err).WithCode("NAME")
	}

	entity.Balance, err = vo.NumberFromRaw(r.Balance)
	if err != nil {
		return RootEntity{}, errors.Wrap(err).WithCode("BALANCE")
	}

	entity.Currency, err = vo.CurrencyFromRaw(r.Currency)
	if err != nil {
		return RootEntity{}, errors.Wrap(err).WithCode("CURRENCY")
	}

	entity.Active, err = vo.BoolFromRaw(r.Active)
	if err != nil {
		return RootEntity{}, errors.Wrap(err).WithCode("ACTIVE")
	}

	return entity, err
}

type RootEntity struct {
	ID       vo.UUID
	Name     vo.String
	Balance  vo.Number
	Currency vo.Currency
	Active   vo.Boolean
}

func (e RootEntity) Raw() RootRaw {
	return RootRaw{
		ID:       e.ID.Raw(),
		Name:     e.Name.Raw(),
		Balance:  e.Balance.Raw(),
		Currency: e.Currency.Raw(),
		Active:   e.Active.Raw(),
	}
}
