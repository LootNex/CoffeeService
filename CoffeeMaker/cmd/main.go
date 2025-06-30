package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LootNex/CoffeeService/CoffeeMaker/configs"
	internalKafka "github.com/LootNex/CoffeeService/CoffeeMaker/internal/kafka"
	kafka "github.com/segmentio/kafka-go"
)

func main() {
	config, err := configs.ConfigLoad()
	if err != nil {
		log.Fatalf("cannot load config err: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		log.Printf("Shutdown signal received: %v", sig)
		cancel()
		time.Sleep(1 * time.Second)
	}()

	reader := internalKafka.NewKafkaReader(config.Kafka.Brokers, config.Kafka.Topic)

	handle := func(msg kafka.Message) {
		log.Printf("received message: value=%s", string(msg.Value))
	}

	go func() {
		err := internalKafka.KafkaConsumer(ctx, reader, handle)
		if err != nil && ctx.Err() == nil {
			log.Printf("cannot read from kafka err:=%v", err)
		}
	}()

	<-ctx.Done()

	log.Println("Exiting CoffeeMaker")
}
