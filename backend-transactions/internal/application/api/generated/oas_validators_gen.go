// Code generated by ogen, DO NOT EDIT.

package generated

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s Currency) Validate() error {
	switch s {
	case "ARS":
		return nil
	case "USD":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *Transaction) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.State.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "state",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.Amount)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "amount",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Currency.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "currency",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s TransactionState) Validate() error {
	switch s {
	case "PENDING":
		return nil
	case "COMPLETED":
		return nil
	case "REJECTED":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *TransactionsGetOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if s.Data == nil {
			return errors.New("nil is invalid value")
		}
		var failures []validate.FieldError
		for i, elem := range s.Data {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "data",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.Count)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "count",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.Skip)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "skip",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.Limit)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "limit",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s TransactionsGetOrderBy) Validate() error {
	switch s {
	case "amount":
		return nil
	case "created_at":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s TransactionsGetOrderDir) Validate() error {
	switch s {
	case "asc":
		return nil
	case "desc":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s TransactionsGetType) Validate() error {
	switch s {
	case "INGRESS":
		return nil
	case "EGRESS":
		return nil
	case "ALL":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *TransactionsTransactionIDPutReq) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.Amount)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "amount",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Currency.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "currency",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}