package null

import (
	"bytes"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

var jsonNilBytes, _ = json.Marshal(nil)

type Null[T any] struct {
	valid bool
	value T
}

func (n Null[T]) IsNull() bool {
	return !n.valid
}

func (n Null[T]) IsNotNull() bool {
	return n.valid
}

func (n Null[T]) Value() T {
	return n.value
}

func (n Null[T]) Equal(other Null[T], cmp func(a, b T) bool) bool {
	if n.IsNotNull() && other.IsNotNull() {
		return cmp(n.Value(), other.Value())
	}
	return n.IsNull() == other.IsNull()
}

func (n Null[T]) MarshalJSON() ([]byte, error) {
	if !n.valid {
		return jsonNilBytes, nil
	}
	return json.Marshal(n.value)
}

func (n *Null[T]) UnmarshalJSON(value []byte) error {
	if bytes.Equal(value, jsonNilBytes) {
		n.valid = false
		return nil
	}
	err := json.Unmarshal(value, &n.value)
	if err != nil {
		return err
	}
	n.valid = true
	return nil
}

func (n Null[T]) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if !n.valid {
		return bson.TypeNull, []byte{}, nil
	}

	return bson.MarshalValue(n.value)
}

func (n *Null[T]) UnmarshalBSONValue(t bsontype.Type, b []byte) error {
	if t == bson.TypeNull {
		n.valid = false
		return nil
	}

	err := bson.UnmarshalValue(t, b, &n.value)
	if err != nil {
		return err
	}

	n.valid = true

	return nil
}

func WithValue[T any](value T) Null[T] {
	return Null[T]{valid: true, value: value}
}

func WithoutValue[T any]() Null[T] {
	return Null[T]{valid: false}
}

func WithError[A any, B any](from Null[A], transformer func(A) (B, error)) (Null[B], error) {
	if from.IsNull() {
		return Null[B]{valid: false}, nil
	}

	value, err := transformer(from.value)
	if err != nil {
		return Null[B]{}, err
	}

	return Null[B]{valid: true, value: value}, nil
}

func WithoutError[A any, B any](from Null[A], transformer func(A) B) Null[B] {
	result, _ := WithError(from, func(a A) (B, error) {
		return transformer(a), nil
	})
	return result
}

func Raw[A interface{ Raw() B }, B any](from Null[A]) Null[B] {
	return WithoutError(from, func(a A) B { return a.Raw() })
}

func Entity[A interface{ Entity() (B, error) }, B any](from Null[A]) (Null[B], error) {
	return WithError(from, func(a A) (B, error) { return a.Entity() })
}

type Equatable[T any] interface {
	Equal(T) bool
}

func Equal[T Equatable[T]](a, b Null[T]) bool {
	return a.Equal(b, func(a, b T) bool { return a.Equal(b) })
}

func EqualPrimitive[T comparable](a, b Null[T]) bool {
	return a.Equal(b, func(a, b T) bool { return a == b })
}
