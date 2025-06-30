package service

import (
	"errors"

	"github.com/LootNex/CoffeeService/ClientManager/internal/kafka"
	"github.com/LootNex/CoffeeService/ClientManager/internal/models"
)

type OrderService struct {
	Service kafka.OrderProducer
}

type OrderCreator interface {
	CreateOrder(order models.Order) error
}

func NewOrderService(p kafka.OrderProducer) OrderService {
	return OrderService{
		Service: p,
	}
}

func (s OrderService) CreateOrder(order models.Order) error {
	if order.Item == "" {
		return errors.New("field item is empty")
	}

	return s.Service.Send(order)
}
