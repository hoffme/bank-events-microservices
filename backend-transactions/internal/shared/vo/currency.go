package vo

import "github.com/hoffme/backend-transactions/internal/shared/errors"

// implementation

type Currency struct {
	raw string
}

func (t Currency) Equal(other Currency) bool {
	return t.raw == other.Raw()
}

func (t Currency) Raw() string {
	return t.raw
}

func CurrencyARS() Currency {
	return Currency{raw: "ARS"}
}
func CurrencyUSD() Currency {
	return Currency{raw: "USD"}
}

// builders

var ErrorCurrencyInvalid = errors.New("VO:CURRENCY:INVALID").WithStatus(400)

func CurrencyFromRaw(value string) (Currency, error) {
	switch value {
	case CurrencyARS().Raw():
		return CurrencyARS(), nil
	case CurrencyUSD().Raw():
		return CurrencyUSD(), nil
	}
	return Currency{}, ErrorCurrencyInvalid
}
