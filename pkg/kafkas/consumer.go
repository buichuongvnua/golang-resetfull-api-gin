package kafkas

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

type Consumer[T comparable] struct {
	Reader *kafka.Reader
	Dialer *kafka.Dialer
	Topic  string
}

func (c *Consumer[T]) CreateConnection(brokers []string, partition int) {
	c.Reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     c.Topic,
		Partition: partition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		MaxWait:   time.Millisecond * 10,
		Dialer:    c.Dialer,
	})
	offset := c.Reader.Offset()
	fmt.Printf("Error while getting offsets: %d\n", offset)
	c.Reader.SetOffset(offset)
}

func (c *Consumer[T]) Read(model T, callback func(T, error)) {
	for {
		message, err := c.Reader.ReadMessage(context.Background())
		if err != nil {
			callback(model, err)
			return
		}
		err = json.Unmarshal(message.Value, &model)
		if err != nil {
			callback(model, err)
			continue
		}
		callback(model, nil)
	}
}
