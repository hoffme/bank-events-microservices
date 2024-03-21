package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hoffme/backend-transactions/internal/shared/events"
	"github.com/hoffme/backend-transactions/internal/shared/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

var _ events.Bus = (*BusRabbitMQ)(nil)

type EventBusRabbitMQConfig struct {
	Url      string `json:"url"`
	Exchange string `json:"exchange"`
}

type BusRabbitMQ struct {
	config    EventBusRabbitMQConfig
	conn      *connection
	consumers map[string]*consumer
}

func NewBusRabbitMQ(config EventBusRabbitMQConfig) (*BusRabbitMQ, error) {
	var err error

	bus := &BusRabbitMQ{
		config:    config,
		consumers: map[string]*consumer{},
	}

	bus.conn, err = newConnection(config.Url)
	if err != nil {
		return nil, err
	}

	err = bus.conn.channel.ExchangeDeclare(
		bus.config.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	err = bus.conn.channel.ExchangeDeclare(
		fmt.Sprintf("%s_DLX", bus.config.Exchange),
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return bus, nil
}

func (b *BusRabbitMQ) Dispatch(ctx context.Context, messages ...events.MessageRaw) {
	nextCtx := context.Background()

	for _, msg := range messages {
		go b.dispatchOne(nextCtx, msg)
	}
}

func (b *BusRabbitMQ) dispatchOne(ctx context.Context, msg events.MessageRaw) {
	defer func() {
		errRecover := recover()
		if errRecover != nil {
			logger.Fatal("dispatch rabbit: %s", errRecover)
		}
	}()

	body, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	err = b.conn.channel.PublishWithContext(
		ctx,
		b.config.Exchange,
		msg.Header.Topic,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			Timestamp:    time.Now(),
			ContentType:  "text/plain",
			MessageId:    msg.Header.ID,
			Body:         body,
		},
	)
	if err != nil {
		panic(err)
	}
}

func (b *BusRabbitMQ) Subscribe(ctx context.Context, handlers ...events.Handler) error {
	for _, handler := range handlers {
		cms, err := newConsumer(b.conn, b.config.Exchange, handler)
		if err != nil {
			return err
		}

		go cms.run()
		b.consumers[handler.Header().ID] = cms
	}

	return nil
}

func (b *BusRabbitMQ) Unsubscribe(ctx context.Context, handlers ...events.Handler) error {
	for _, handler := range handlers {
		cms, ok := b.consumers[handler.Header().ID]
		if !ok {
			continue
		}

		cms.close()
		delete(b.consumers, handler.Header().ID)
	}

	return nil
}

func (b *BusRabbitMQ) Close() {
	err := b.conn.close()
	if err != nil {
		logger.Warn("rabbit closed: %s", err.Error())
	}
}

// consumer

type consumer struct {
	handler  events.Handler
	conn     *connection
	exchange string
	consumer <-chan amqp.Delivery
}

func newConsumer(conn *connection, exchange string, handler events.Handler) (*consumer, error) {
	csm := &consumer{
		conn:     conn,
		exchange: exchange,
		handler:  handler,
	}

	err := csm.declareQueue()
	if err != nil {
		return nil, err
	}

	err = csm.declareRetry()
	if err != nil {
		return nil, err
	}

	err = csm.declareConsume()
	if err != nil {
		return nil, err
	}

	return csm, nil
}

func (c *consumer) declareQueue() error {
	_, err := c.conn.channel.QueueDeclare(
		c.handler.Header().QueueTopic,
		true,
		false,
		false,
		true,
		amqp.Table{
			"x-dead-letter-exchange": fmt.Sprintf("%s_DLX", c.exchange),
		},
	)
	if err != nil {
		return err
	}

	for _, consumeTopic := range c.handler.Header().ConsumeTopics {
		err = c.conn.channel.QueueBind(
			c.handler.Header().QueueTopic,
			consumeTopic,
			c.exchange,
			false,
			nil,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *consumer) declareRetry() error {
	if c.handler.Header().TTL < 0 {
		return nil
	}

	var ttl int64 = 2000
	if c.handler.Header().TTL > 0 {
		ttl = c.handler.Header().TTL
	}

	queueName := fmt.Sprintf("%s_DLQ", c.handler.Header().QueueTopic)

	_, err := c.conn.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		true,
		amqp.Table{
			"x-dead-letter-exchange": c.exchange,
			"x-message-ttl":          ttl,
		},
	)
	if err != nil {
		return err
	}

	for _, consumeTopic := range c.handler.Header().ConsumeTopics {
		err = c.conn.channel.QueueBind(
			queueName,
			consumeTopic,
			fmt.Sprintf("%s_DLX", c.exchange),
			false,
			nil,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *consumer) declareConsume() error {
	var err error

	c.consumer, err = c.conn.channel.Consume(
		c.handler.Header().QueueTopic,
		c.handler.Header().ID,
		false,
		false,
		false,
		false,
		nil,
	)

	return err
}

func (c *consumer) run() {
	for msg := range c.consumer {
		logPrefix := fmt.Sprintf("queue '%s' consuming '%s'", c.handler.Header().QueueTopic, msg.MessageId)

		chanErr := make(chan interface{}, 1)

		go c.execute(msg, chanErr)

		err := <-chanErr
		if err != nil {
			logger.Warn("%s: %s", logPrefix, err)

			retry := c.retry(msg)

			if retry {
				err = msg.Nack(false, false)
				if err != nil {
					logger.Error("%s with error nack: %s", logPrefix, err)
				}
			} else {
				err = msg.Ack(false)
				if err != nil {
					logger.Error("%s with error ack: %s", logPrefix, err)
				}
			}

			go c.sendFailure(msg, err, retry)

			msg.MessageCount += 1
		} else {
			err = msg.Ack(false)
			if err != nil {
				logger.Error("%s with success ack: %s", logPrefix, err)
			}
		}
	}
}

func (c *consumer) close() {}

func (c *consumer) execute(delivery amqp.Delivery, chanErr chan interface{}) {
	defer func() {
		chanErr <- recover()
	}()

	ctx := context.Background()

	msg := events.MessageRaw{}
	err := json.Unmarshal(delivery.Body, &msg)
	if err != nil {
		panic(err)
	}

	err = c.handler.Resolve(ctx, msg)
	if err != nil {
		panic(err)
	}
}

func (c *consumer) retry(delivery amqp.Delivery) bool {
	xDeath, ok := delivery.Headers["x-death"].([]interface{})
	if !ok {
		return true
	}

	count, ok := xDeath[0].(amqp.Table)["count"].(int64)
	if !ok {
		return false
	}

	return count < c.handler.Header().MaxIntents
}

func (c *consumer) sendFailure(delivery amqp.Delivery, errRecover any, retry bool) {
	msg := fmt.Sprintf("%s", errRecover)
	if err, ok := errRecover.(error); ok {
		msg = err.Error()
	}
	logger.Error("RABBIT_CONSUME_FAILURE", delivery.MessageId, msg)
}

// connection

type connection struct {
	url        string
	connection *amqp.Connection
	channel    *amqp.Channel
}

func newConnection(url string) (*connection, error) {
	var err error

	conn := &connection{url: url}

	conn.connection, err = amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	conn.channel, err = conn.connection.Channel()
	if err != nil {
		errClose := conn.close()
		if errClose != nil {
			logger.Warn("%s", errClose.Error())
		}

		return nil, err
	}

	return conn, nil
}

func (c *connection) close() error {
	err := c.channel.Close()
	if err != nil {
		logger.Warn(err.Error())
	}

	err = c.connection.Close()
	if err != nil {
		logger.Warn(err.Error())
	}

	return nil
}
