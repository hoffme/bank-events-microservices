package infrastructure

import (
	"github.com/hoffme/backend-transactions/internal/application"
)

func (i Infrastructure) Ports() application.Ports {
	return application.Ports{
		APIServer: i.server,
		EventBus:  i.rabbit,
	}
}
