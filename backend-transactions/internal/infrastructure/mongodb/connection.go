package mongodb

import (
	"context"

	"github.com/hoffme/backend-transactions/internal/shared/logger"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Uri string `json:"uri"`
}

type Connection struct {
	config Config
	client *mongo.Client
}

func Connect(ctx context.Context, config Config) (*Connection, error) {
	opts := options.Client().ApplyURI(config.Uri).SetMonitor(&event.CommandMonitor{})

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &Connection{
		config: config,
		client: client,
	}, nil
}

func (c *Connection) Close(ctx context.Context) {
	err := c.client.Disconnect(ctx)
	if err != nil {
		logger.Warn("close mongo %w", err)
	}
}
