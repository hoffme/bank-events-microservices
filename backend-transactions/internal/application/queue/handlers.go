package queue

import (
	"github.com/hoffme/backend-transactions/internal/shared/events"

	"github.com/hoffme/backend-transactions/internal/domain"

	"github.com/hoffme/backend-transactions/internal/application/queue/handlers"
)

func Handlers(modules domain.Dependencies) ([]events.Handler, error) {
	result := []events.Handler{
		handlers.NewAccountCreate(modules),
		handlers.NewAccountUpdateName(modules),
		handlers.NewAccountInactivate(modules),
		handlers.NewAccountActivate(modules),
		handlers.NewTransactionAuthorize(modules),
	}

	return result, nil
}
