package main

import (
	"log"

	"github.com/LootNex/CoffeeService/CoffeeMaker/internal/kafka"
)

func main() {

	err := kafka.KafkaConsumer([]string{"kafka:9092"}, "orders")

	if err != nil {
		log.Printf("cannot read from kafka err:=%v", err)
	}

}
