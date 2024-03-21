package infrastructure

import (
	"context"

	"github.com/hoffme/backend-transactions/internal/shared/env"
	"github.com/hoffme/backend-transactions/internal/shared/server"

	"github.com/hoffme/backend-transactions/internal/infrastructure/http"
	"github.com/hoffme/backend-transactions/internal/infrastructure/mongodb"
	"github.com/hoffme/backend-transactions/internal/infrastructure/rabbitmq"
)

type Infrastructure struct {
	rabbit *rabbitmq.BusRabbitMQ
	mongo  *mongodb.Connection
	server *http.Server
}

func Initialize(ctx context.Context) (Infrastructure, error) {
	var err error

	result := Infrastructure{}

	// Storage

	result.mongo, err = mongodb.Connect(ctx, mongodb.Config{
		Uri: env.String("MONGO_URI").Value(),
	})
	if err != nil {
		return Infrastructure{}, err
	}

	// Event Bus

	result.rabbit, err = rabbitmq.NewBusRabbitMQ(rabbitmq.EventBusRabbitMQConfig{
		Url:      env.String("RABBIT_URL").Value(),
		Exchange: env.String("RABBIT_EXCHANGE").Value(),
	})
	if err != nil {
		return Infrastructure{}, err
	}

	// Servers

	result.server = http.New(http.Config{
		Addr: env.String("HOST").Value(),
	})

	return result, nil
}

func (i Infrastructure) Close(ctx context.Context) {
	i.server.Close()
	i.rabbit.Close()
	i.mongo.Close(ctx)
}

func (i Infrastructure) Servers() []server.Server {
	return []server.Server{
		i.server,
	}
}
