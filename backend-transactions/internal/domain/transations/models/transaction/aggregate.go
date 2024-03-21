package transaction

import (
	"github.com/hoffme/backend-transactions/internal/shared/errors"
	"github.com/hoffme/backend-transactions/internal/shared/events"
)

type Aggregate struct {
	root RootEntity

	events *events.Recorder
}

func (a *Aggregate) Raw() Raw {
	return Raw{
		Root: a.root.Raw(),
	}
}

type Raw struct {
	Root RootRaw
}

func (r Raw) Instance() (*Aggregate, error) {
	var err error

	instance := &Aggregate{}

	instance.root, err = r.Root.Entity()
	if err != nil {
		return nil, errors.Wrap(err).WithCode(Error.Code())
	}

	instance.events = events.NewRecorder()

	return instance, err
}
