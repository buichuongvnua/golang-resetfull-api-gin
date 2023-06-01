package kafkas

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type Publisher struct {
	Writer *kafka.Writer
}

func NewPublisher(brokers []string) (*Publisher, error) {
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      brokers,
		BatchBytes:   1000000,
		BatchTimeout: time.Millisecond * 5,
		Balancer:     &kafka.RoundRobin{},
		Dialer:       dialer,
	})

	return &Publisher{
		Writer: writer,
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
