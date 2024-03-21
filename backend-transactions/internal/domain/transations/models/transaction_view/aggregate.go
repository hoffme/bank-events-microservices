package transaction_view

import (
	"github.com/hoffme/backend-transactions/internal/shared/errors"
)

type Aggregate struct {
	root RootEntity
	from AccountEntity
	to   AccountEntity
}

func (a *Aggregate) Raw() Raw {
	return Raw{
		Root: a.root.Raw(),
		From: a.from.Raw(),
		To:   a.to.Raw(),
	}
}

type Raw struct {
	Root RootRaw
	From AccountRaw
	To   AccountRaw
}

func (r Raw) Instance() (*Aggregate, error) {
	var err error

	instance := &Aggregate{}

	instance.root, err = r.Root.Entity()
	if err != nil {
		return nil, errors.Wrap(err).WithCode(Error.Code())
	}

	instance.from, err = r.From.Entity()
	if err != nil {
		return nil, errors.Wrap(err).WithCode(Error.Code()).WithCode("FROM_ACCOUNT")
	}

	instance.to, err = r.To.Entity()
	if err != nil {
		return nil, errors.Wrap(err).WithCode(Error.Code()).WithCode("TO_ACCOUNT")
	}

	return instance, err
}
