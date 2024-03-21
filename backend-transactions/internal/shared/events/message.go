package events

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/hoffme/backend-transactions/internal/shared/logger"

	"github.com/google/uuid"
)

// Topic

func EventTopic(tokens ...string) string {
	return fmt.Sprintf("%s.%s", "app.transactions.evt", strings.Join(tokens, "."))
}

func QueueTopic(tokens ...string) string {
	return fmt.Sprintf("%s.%s", "app.transactions.qeu", strings.Join(tokens, "."))
}

// Payload

type Payload interface {
	Topic() string
}

func PayloadEncode(data Payload) ([]byte, error) {
	return json.Marshal(data)
}

func PayloadDecode[P Payload](data []byte) (result P, err error) {
	err = json.Unmarshal(data, &result)
	return
}

// Params

type MessageHeader struct {
	ID        string            `json:"id"`
	Topic     string            `json:"topic"`
	Timestamp string            `json:"timestamp"`
	Metadata  map[string]string `json:"metadata"`
}

type Message[P Payload] struct {
	Header  MessageHeader `json:"header"`
	Payload P             `json:"payload"`
}

type MessageRaw struct {
	Header  MessageHeader `json:"header"`
	Payload []byte        `json:"payload"`
}

type MessageGeneric = Message[Payload]

func (e Message[P]) ToGeneric() MessageGeneric {
	return Message[Payload]{
		Header:  e.Header,
		Payload: e.Payload,
	}
}

func (e Message[P]) ToRaw() MessageRaw {
	bytes, err := PayloadEncode(e.Payload)
	if err != nil {
		logger.Warn("message encode payload: %s", err.Error())
	}

	return MessageRaw{
		Header:  e.Header,
		Payload: bytes,
	}
}

// Builders

func NewMessage[P Payload](payload P) Message[P] {
	return Message[P]{
		Header: MessageHeader{
			ID:        uuid.NewString(),
			Topic:     payload.Topic(),
			Timestamp: time.Now().Format(time.RFC3339Nano),
			Metadata:  map[string]string{},
		},
		Payload: payload,
	}
}

func NewMessageGeneric(payload Payload) MessageGeneric {
	return NewMessage(payload)
}

func NewMessageRaw(payload Payload) MessageRaw {
	bytes, _ := PayloadEncode(payload)
	return MessageRaw{
		Header: MessageHeader{
			ID:        uuid.NewString(),
			Topic:     payload.Topic(),
			Timestamp: time.Now().Format(time.RFC3339Nano),
			Metadata:  map[string]string{},
		},
		Payload: bytes,
	}
}

func MessageFromRaw[P Payload](raw MessageRaw) (Message[P], error) {
	payload, err := PayloadDecode[P](raw.Payload)
	if err != nil {
		return Message[P]{}, err
	}

	return Message[P]{
		Header:  raw.Header,
		Payload: payload,
	}, nil
}

func MessageFromGeneric[P Payload](event MessageGeneric) (Message[P], bool) {
	data, ok := event.Payload.(P)
	if !ok {
		return Message[P]{}, false
	}

	return Message[P]{
		Header:  event.Header,
		Payload: data,
	}, true
}

// Params Config

type MessageConfig func(message *Message[Payload])

func MessageConfigSetHeader(header MessageHeader) MessageConfig {
	return func(m *Message[Payload]) {
		m.Header = header
	}
}
