package producer

import (
	"context"
	"message-queue/config"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

// Connect to Broker
func ConnectBroker(cfg *config.Config) Producer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{cfg.BROKER_ADDR},
		Topic:   cfg.BROKER_TOPIC,
		Async:   true})

	return Producer{
		writer: writer,
	}
}

type QueuePayload struct {
	ProductID int
}

func (p Producer) Produce(payload []byte) error {
	ctx := context.Background()

	err := p.writer.WriteMessages(ctx, kafka.Message{
		Value: payload,
	})
	if err != nil {
		return err
	}

	return nil
}
