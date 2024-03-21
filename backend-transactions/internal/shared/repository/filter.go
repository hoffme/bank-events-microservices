package repository

import (
	"github.com/hoffme/backend-transactions/internal/shared/null"
)

type Where struct {
	Field    string      `json:"field"`
	Operator WhereOp     `json:"operator"`
	Value    interface{} `json:"value"`
}

type Order struct {
	By  string   `json:"by"`
	Dir OrderDir `json:"dir"`
}

type FindFilter struct {
	Query null.Null[string]  `json:"query,omitempty"`
	Where null.Null[Where]   `json:"where,omitempty"`
	Order null.Null[[]Order] `json:"order,omitempty"`
}

type SearchFilter struct {
	Query null.Null[string]  `json:"query,omitempty"`
	Where null.Null[Where]   `json:"where,omitempty"`
	Limit null.Null[int64]   `json:"limit,omitempty"`
	Skip  null.Null[int64]   `json:"skip,omitempty"`
	Order null.Null[[]Order] `json:"order,omitempty"`
}

type SearchResult[T any] struct {
	Data  []T   `json:"data"`
	Count int64 `json:"count"`
	Limit int64 `json:"limit"`
	Skip  int64 `json:"skip"`
}

func MapResult[T any, B any](params SearchResult[T], parser func(T) B) SearchResult[B] {
	result := SearchResult[B]{
		Data:  make([]B, len(params.Data)),
		Count: params.Count,
		Limit: params.Limit,
		Skip:  params.Skip,
	}

	for i, row := range params.Data {
		result.Data[i] = parser(row)
	}

	return result
}

func MapResultWithError[T any, B any](params SearchResult[T], parser func(T) (B, error)) (SearchResult[B], error) {
	var err error

	result := SearchResult[B]{
		Data:  make([]B, len(params.Data)),
		Count: params.Count,
		Limit: params.Limit,
		Skip:  params.Skip,
	}

	for i, row := range params.Data {
		result.Data[i], err = parser(row)
		if err != nil {
			return SearchResult[B]{}, err
		}
	}

	return result, nil
}
