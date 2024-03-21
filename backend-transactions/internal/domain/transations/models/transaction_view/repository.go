package transaction_view

import (
	"context"

	"github.com/hoffme/backend-transactions/internal/shared/repository"
)

type Repository interface {
	Get(ctx context.Context, id string) (*Aggregate, error)
	Search(ctx context.Context, filter repository.SearchFilter) (repository.SearchResult[*Aggregate], error)
}
