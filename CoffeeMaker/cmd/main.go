package main

import (
	"log"
	"time"

	internalKafka "github.com/LootNex/CoffeeService/CoffeeMaker/internal/kafka"
	kafka "github.com/segmentio/kafka-go"
)

func main() {

	reader := &internalKafka.ReadKafkaReader{Reader: kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"kafka:9092"},
		GroupID:  "coffee-prep-service",
		Topic:    "orders",
		MinBytes: 10e3,
		MaxBytes: 10e6,
		MaxWait:  1 * time.Second,
	}),
	}

	handle := func(msg kafka.Message) {
		log.Printf("received message: value=%s", string(msg.Value))
	}

	err := internalKafka.KafkaConsumer(reader, handle)

	if err != nil {
		log.Printf("cannot read from kafka err:=%v", err)
	}

}
