package main

import (
	"context"
	"sync"

	"github.com/hoffme/backend-transactions/internal/application"
	"github.com/hoffme/backend-transactions/internal/infrastructure"
	"github.com/hoffme/backend-transactions/internal/migrations"
	"github.com/hoffme/backend-transactions/internal/shared/logger"
)

func main() {
	migrations.Migrate()

	ctx := context.Background()

	// Infrastructure

	infra, err := infrastructure.Initialize(ctx)
	if err != nil {
		panic(err)
	}
	defer infra.Close(ctx)

	// Handlers

	err = application.Setup(ctx, infra.Ports(), infra.DomainDependencies())
	if err != nil {
		panic(err)
	}

	// Run

	wg := sync.WaitGroup{}

	for _, server := range infra.Servers() {
		wg.Add(1)
		go server.Run(&wg)
	}

	logger.Info("🚀 Started Trank")

	wg.Wait()
}
