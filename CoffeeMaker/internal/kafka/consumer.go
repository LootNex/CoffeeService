package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func KafkaConsumer(brokers []string, topic string) error {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  "coffee-prep-service",
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
		MaxWait:  1 * time.Second,
	})

	defer reader.Close()

	log.Println("waiting for messages...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("error reading message:", err)
		}
		log.Printf("received message: value=%s", string(msg.Value))
	}
}
