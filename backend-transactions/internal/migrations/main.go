package migrations

import (
	"context"

	"github.com/hoffme/backend-transactions/internal/shared/env"
	"github.com/hoffme/backend-transactions/internal/shared/logger"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var runners = map[string]func(context.Context, *mongo.Client) error{
	"initial": run_initial,
}

func Migrate() {
	mongoUri := env.String("MONGO_URI").Value()

	ctx := context.Background()

	opts := options.Client().ApplyURI(mongoUri).SetMonitor(&event.CommandMonitor{})

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	for name, runner := range runners {
		logger.Debug("migrate: %s", name)
		err = runner(ctx, client)
		if err != nil {
			panic(err)
		}
		logger.Debug("OK")
	}
}
