package vo

import "github.com/hoffme/backend-transactions/internal/shared/errors"

// implementation

type TransactionState struct {
	raw string
}

func (t TransactionState) Equal(other TransactionState) bool {
	return t.raw == other.Raw()
}

func (t TransactionState) Raw() string {
	return t.raw
}

func TransactionStatePending() TransactionState {
	return TransactionState{raw: "PENDING"}
}
func TransactionStateCompleted() TransactionState {
	return TransactionState{raw: "COMPLETED"}
}
func TransactionStateRejected() TransactionState {
	return TransactionState{raw: "REJECTED"}
}

// builders

var ErrorTransactionStateInvalid = errors.New("VO:TRANSACTION_STATE:INVALID").WithStatus(400)

func TransactionStateFromRaw(value string) (TransactionState, error) {
	switch value {
	case TransactionStatePending().Raw():
		return TransactionStatePending(), nil
	case TransactionStateCompleted().Raw():
		return TransactionStateCompleted(), nil
	case TransactionStateRejected().Raw():
		return TransactionStateRejected(), nil
	}
	return TransactionState{}, ErrorCurrencyInvalid
}
