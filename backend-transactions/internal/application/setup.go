package application

import (
	"context"

	"github.com/hoffme/backend-transactions/internal/shared/events"
	"github.com/hoffme/backend-transactions/internal/shared/server"

	"github.com/hoffme/backend-transactions/internal/domain"

	"github.com/hoffme/backend-transactions/internal/application/api"
	"github.com/hoffme/backend-transactions/internal/application/queue"
)

type Ports struct {
	APIServer server.Server
	EventBus  events.Bus
}

func Setup(ctx context.Context, ports Ports, deps domain.Dependencies) error {
	apiHandler, err := api.Handler(deps)
	if err != nil {
		return err
	}

	eventsHandlers, err := queue.Handlers(deps)
	if err != nil {
		return err
	}

	err = ports.EventBus.Subscribe(ctx, eventsHandlers...)
	if err != nil {
		return err
	}

	ports.APIServer.SetHandler(apiHandler)

	return nil
}
