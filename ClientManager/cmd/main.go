package main

import (
	"log"
	"net/http"

	"github.com/LootNex/CoffeeService/ClientManager/internal/handler"
	"github.com/LootNex/CoffeeService/ClientManager/internal/kafka"
	"github.com/LootNex/CoffeeService/ClientManager/internal/service"
)

func main() {
	producer := kafka.NewKafkaProducer([]string{"kafka:9092"}, "orders")
	defer producer.Writer.Close()

	OrderService := service.NewOrderService(producer)

	handler := handler.Newhandler(&OrderService)

	http.HandleFunc("/order", handler.CreateOrderHandler)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("cannot start http server")
	}

}
