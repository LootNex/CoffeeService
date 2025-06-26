package kafka

import (
	"context"
	"log"

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

func KafkaConsumer(reader KafkaReader, handle func(kafka.Message)) error {

	log.Println("Kafka consumer is ready")

	defer reader.Close()

	for {
		log.Println("waiting for messages...")
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error reading message: %v", err)
			return err
		}
		handle(msg)
	}
}
