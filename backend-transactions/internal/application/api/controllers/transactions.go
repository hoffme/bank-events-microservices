package controllers

import (
	"context"
	"time"

	"github.com/hoffme/backend-transactions/internal/shared/null"
	"github.com/hoffme/backend-transactions/internal/shared/repository"
	"github.com/hoffme/backend-transactions/internal/shared/vo"

	"github.com/hoffme/backend-transactions/internal/domain/transations/models/transaction_view"
	"github.com/hoffme/backend-transactions/internal/domain/transations/services/transactions"

	"github.com/hoffme/backend-transactions/internal/application/api/generated"
)

var mapState = map[string]generated.TransactionState{
	vo.TransactionStatePending().Raw():   generated.TransactionStatePENDING,
	vo.TransactionStateCompleted().Raw(): generated.TransactionStateCOMPLETED,
	vo.TransactionStateRejected().Raw():  generated.TransactionStateREJECTED,
}
var mapCurrencyDecode = map[string]generated.Currency{
	vo.CurrencyARS().Raw(): generated.CurrencyARS,
	vo.CurrencyUSD().Raw(): generated.CurrencyUSD,
}
var mapCurrencyEncode = map[generated.Currency]vo.Currency{
	generated.CurrencyARS: vo.CurrencyARS(),
	generated.CurrencyUSD: vo.CurrencyUSD(),
}

func (c Controllers) TransactionsGet(ctx context.Context, params generated.TransactionsGetParams) (*generated.TransactionsGetOK, error) {
	whereAnd := make([]repository.Where, 0)
	var limit int64 = 10
	var skip int64 = 0
	orderBy := "created_at"
	orderDir := repository.OrderDirDesc

	if params.AccountID.IsSet() {
		if params.Type.Value == generated.TransactionsGetTypeINGRESS {
			whereAnd = append(whereAnd, repository.Where{
				Field:    "to_account.id",
				Operator: repository.WhereOpEq,
				Value:    params.AccountID.Value,
			})
		} else if params.Type.Value == generated.TransactionsGetTypeEGRESS {
			whereAnd = append(whereAnd, repository.Where{
				Field:    "from_account.id",
				Operator: repository.WhereOpEq,
				Value:    params.AccountID.Value,
			})
		} else {
			whereAnd = append(whereAnd, repository.Where{
				Operator: repository.WhereOpOr,
				Value: []repository.Where{
					{Field: "from_account.id", Operator: repository.WhereOpEq, Value: params.AccountID.Value},
					{Field: "to_account.id", Operator: repository.WhereOpEq, Value: params.AccountID.Value},
				},
			})
		}
	}
	if params.DateFrom.IsSet() {
		whereAnd = append(whereAnd, repository.Where{
			Field:    "created_at",
			Operator: repository.WhereOpGte,
			Value:    params.DateFrom,
		})
	}
	if params.DateTo.IsSet() {
		whereAnd = append(whereAnd, repository.Where{
			Field:    "created_at",
			Operator: repository.WhereOpLte,
			Value:    params.DateFrom,
		})
	}
	if len(params.State) > 0 {
		stateMap := map[generated.TransactionState]vo.TransactionState{
			generated.TransactionStatePENDING:   vo.TransactionStatePending(),
			generated.TransactionStateREJECTED:  vo.TransactionStateRejected(),
			generated.TransactionStateCOMPLETED: vo.TransactionStateCompleted(),
		}

		states := make([]string, len(params.State))
		for i, state := range params.State {
			states[i] = stateMap[state].Raw()
		}

		whereAnd = append(whereAnd, repository.Where{
			Field:    "state",
			Operator: repository.WhereOpIn,
			Value:    states,
		})
	}
	if len(params.Currency) > 0 {
		currencyMap := map[generated.Currency]vo.Currency{
			generated.CurrencyARS: vo.CurrencyARS(),
			generated.CurrencyUSD: vo.CurrencyUSD(),
		}

		currencies := make([]string, len(params.Currency))
		for i, currency := range params.Currency {
			currencies[i] = currencyMap[currency].Raw()
		}

		whereAnd = append(whereAnd, repository.Where{
			Field:    "currency",
			Operator: repository.WhereOpIn,
			Value:    currencies,
		})
	}
	if params.Limit.Set {
		limit = int64(params.Limit.Value)
	}
	if params.Skip.Set {
		skip = int64(params.Skip.Value)
	}
	if params.OrderBy.Set {
		mapOrderBy := map[generated.TransactionsGetOrderBy]string{
			generated.TransactionsGetOrderByAmount:    "amount",
			generated.TransactionsGetOrderByCreatedAt: "created_at",
		}
		orderBy = mapOrderBy[params.OrderBy.Value]
	}
	if params.OrderDir.Set {
		mapOrderDir := map[generated.TransactionsGetOrderDir]repository.OrderDir{
			generated.TransactionsGetOrderDirAsc:  repository.OrderDirAsc,
			generated.TransactionsGetOrderDirDesc: repository.OrderDirDesc,
		}
		orderDir = mapOrderDir[params.OrderDir.Value]
	}

	filter := repository.SearchFilter{
		Query: null.WithoutValue[string](),
		Where: null.WithValue[repository.Where](repository.Where{Operator: repository.WhereOpAnd, Value: whereAnd}),
		Limit: null.WithValue[int64](limit),
		Skip:  null.WithValue[int64](skip),
		Order: null.WithValue([]repository.Order{{By: orderBy, Dir: orderDir}}),
	}

	transactions, err := c.Deps.Repositories.TransactionView.Search(ctx, filter)
	if err != nil {
		return nil, err
	}

	result := generated.TransactionsGetOK{
		Data:  make([]generated.Transaction, len(transactions.Data)),
		Count: float64(transactions.Count),
		Skip:  float64(transactions.Skip),
		Limit: float64(transactions.Limit),
	}

	for i, row := range transactions.Data {
		result.Data[i] = c.transactionToGenerated(row)
	}

	return &result, nil
}

func (c Controllers) TransactionsTransactionIDGet(ctx context.Context, params generated.TransactionsTransactionIDGetParams) (*generated.Transaction, error) {
	instance, err := c.Deps.Repositories.TransactionView.Get(ctx, params.TransactionID)
	if err != nil {
		return nil, err
	}

	result := c.transactionToGenerated(instance)

	return &result, nil
}

func (c Controllers) TransactionsTransactionIDPut(ctx context.Context, req generated.OptTransactionsTransactionIDPutReq, params generated.TransactionsTransactionIDPutParams) (*generated.TransactionsTransactionIDPutCreated, error) {
	transferParams := transactions.TransferParamsRaw{
		TransferID:    params.TransactionID,
		FromAccountID: req.Value.FromAccountID,
		ToAccountID:   req.Value.ToAccountID,
		Currency:      mapCurrencyEncode[req.Value.Currency].Raw(),
		Amount:        int64(req.Value.Amount),
	}

	err := c.Deps.Services.Transactions.Transfer(ctx, transferParams)
	if err != nil {
		return nil, err
	}

	return &generated.TransactionsTransactionIDPutCreated{Ok: true}, nil
}

func (c Controllers) transactionToGenerated(value *transaction_view.Aggregate) generated.Transaction {
	return generated.Transaction{
		ID:    value.GetID(),
		State: mapState[value.GetState()],
		From: generated.TransactionFrom{
			ID:   value.GetFromAccountID(),
			Name: value.GetFromAccountName(),
		},
		To: generated.TransactionTo{
			ID:   value.GetToAccountID(),
			Name: value.GetToAccountName(),
		},
		Amount:    float64(value.GetAmount()),
		Currency:  mapCurrencyDecode[value.GetCurrency()],
		CreatedAt: value.GetCreatedAt().Format(time.RFC3339),
	}
}
