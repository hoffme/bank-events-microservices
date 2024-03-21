package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Parser[T any] func(string) (T, error)

type Var[T any] struct {
	key          string
	required     bool
	defaultValue T
	parser       Parser[T]
}

func (v Var[T]) SetRequired(value bool) Var[T] {
	v.required = value
	return v
}

func (v Var[T]) SetDefaultValue(value T) Var[T] {
	v.defaultValue = value
	return v
}

func (v Var[T]) SetParser(parser Parser[T]) Var[T] {
	v.parser = parser
	return v
}

func (v Var[T]) Value() T {
	stringValue, ok := os.LookupEnv(v.key)
	if !ok {
		if v.required {
			panic(fmt.Errorf("missing required enviroment variable '%s'", v.key))
		}

		return v.defaultValue
	}

	result, err := v.parser(stringValue)
	if err != nil {
		panic(fmt.Errorf("cannot parse enviroment variable '%s': '%s'", v.key, stringValue))
	}

	return result
}

func Get[T any](key string, parser Parser[T]) Var[T] {
	return Var[T]{
		key:    key,
		parser: parser,
	}
}

func Bool(key string) Var[bool] {
	return Get(key, ParseBool)
}

func String(key string) Var[string] {
	return Get(key, ParseString)
}

func Int(key string) Var[int] {
	return Get(key, ParseInt)
}

func Float64(key string) Var[float64] {
	return Get(key, ParseFloat64)
}

func Time(key string) Var[time.Time] {
	return Get(key, ParseTime(time.RFC3339))
}

func Duration(key string) Var[time.Duration] {
	return Get(key, ParseDuration)
}

func ParseBool(v string) (bool, error) {
	if strings.ToLower(v) == "true" {
		return true, nil
	}
	if strings.ToLower(v) == "false" {
		return false, nil
	}

	return false, fmt.Errorf("invalid bool value")
}

func ParseString(v string) (string, error) {
	return v, nil
}

func ParseInt(v string) (int, error) {
	return strconv.Atoi(v)
}

func ParseFloat64(v string) (float64, error) {
	return strconv.ParseFloat(v, 64)
}

func ParseTime(layout string) Parser[time.Time] {
	return func(v string) (time.Time, error) {
		return time.Parse(layout, v)
	}
}

func ParseDuration(v string) (time.Duration, error) {
	return time.ParseDuration(v)
}
