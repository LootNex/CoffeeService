package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LootNex/CoffeeService/ClientManager/configs"
	"github.com/LootNex/CoffeeService/ClientManager/internal/handler"
	"github.com/LootNex/CoffeeService/ClientManager/internal/kafka"
	"github.com/LootNex/CoffeeService/ClientManager/internal/service"
)

func main() {

	config, err := configs.ConfigLoad()
	if err != nil {
		log.Fatalf("cannot load config err: %v", err)
	}

	producer := kafka.NewKafkaProducer(config.Kafka.Brokers, config.Kafka.Topic)
	defer producer.Writer.Close()

	OrderService := service.NewOrderService(producer)

	handler := handler.Newhandler(&OrderService)

	http.HandleFunc("/order", handler.CreateOrderHandler)

	addr := fmt.Sprintf(":%s", config.Server.Port)

	log.Println("ClientManager is ready!")

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("cannot start http server")
	}

}
