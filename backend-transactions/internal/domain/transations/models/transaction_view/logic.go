package transaction_view

import (
	"time"

	"github.com/hoffme/backend-transactions/internal/shared/null"
)

func (a *Aggregate) GetID() string {
	return a.root.ID.Raw()
}

func (a *Aggregate) GetFromAccountID() string {
	return a.from.ID.Raw()
}

func (a *Aggregate) GetFromAccountName() string {
	return a.from.Name.Raw()
}

func (a *Aggregate) GetToAccountID() string {
	return a.to.ID.Raw()
}

func (a *Aggregate) GetToAccountName() string {
	return a.to.Name.Raw()
}

func (a *Aggregate) GetState() string {
	return a.root.State.Raw()
}

func (a *Aggregate) GetCurrency() string {
	return a.root.Currency.Raw()
}

func (a *Aggregate) GetAmount() int64 {
	return a.root.Amount.Raw()
}

func (a *Aggregate) GetCreatedAt() time.Time {
	return a.root.CreatedAt.Raw()
}

func (a *Aggregate) GetFinishedAt() null.Null[time.Time] {
	return null.Raw(a.root.FinishedAt)
}

func (a *Aggregate) GetFinished() bool {
	return a.root.FinishedAt.IsNotNull()
}
