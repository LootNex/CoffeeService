package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaReader interface {
	Close() error
	ReadMessage(context.Context) (kafka.Message, error)
}

type ReadKafkaReader struct {
	Reader *kafka.Reader
}

func (r *ReadKafkaReader) Close() error {
	return r.Reader.Close()
}

func (r *ReadKafkaReader) ReadMessage(ctx context.Context) (kafka.Message, error) {
	return r.Reader.ReadMessage(ctx)
}

func KafkaConsumer(ctx context.Context, reader KafkaReader, handle func(kafka.Message)) error {

	log.Println("Kafka consumer is ready")

	defer reader.Close()

	for {
		log.Println("waiting for messages...")
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("error reading message: %v", err)
			return err
		}
		handle(msg)
	}
}

func NewKafkaReader(brokers []string, topic string) *ReadKafkaReader {

	return &ReadKafkaReader{Reader: kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Partition:   0,
		StartOffset: kafka.FirstOffset,
		Topic:       topic,
		MinBytes:    10e3,
		MaxBytes:    10e6,
		MaxWait:     1 * time.Second,
	}),
	}

}
