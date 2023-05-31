package kafkas

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type Publisher struct {
	Writer *kafka.Writer
	Dialer *kafka.Dialer
}

func NewPublisher() (*Publisher, error) {
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{"localhost:9092"},
		BatchBytes:   1000000,
		BatchTimeout: time.Millisecond * 5,
		Balancer:     &kafka.RoundRobin{},
		Dialer:       dialer,
	})

	return &Publisher{
		Writer: writer,
		Dialer: dialer,
	}, nil
}

func (p *Publisher) Publish(ctx context.Context, topic string, key string, value []byte) error {
	err := p.Writer.WriteMessages(context.Background(), kafka.Message{
		Topic:  topic,
		Offset: 0,
		Key:    []byte(key),
		Value:  value,
	})
	p.Writer.Close()
	return err
}
