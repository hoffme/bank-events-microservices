package account

import "context"

type Repository interface {
	Get(ctx context.Context, id string) (*Aggregate, error)
	Save(ctx context.Context, instance *Aggregate) error
}
