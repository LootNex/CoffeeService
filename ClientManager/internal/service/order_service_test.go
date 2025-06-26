package service

import (
	"testing"

	"github.com/LootNex/CoffeeService/ClientManager/pkg/models"
)

type MockSend struct {
	SendFunc func(order models.Order) error
}

func (m MockSend) Send(order models.Order) error {

	return m.SendFunc(order)
}

func TestCreateOrder(t *testing.T) {

	mock := MockSend{
		SendFunc: func(order models.Order) error {
			return nil
		},
	}

	service := NewOrderService(mock)

	order := models.Order{
		ID:    "1",
		Buyer: "N",
		Item:  "latte",
	}

	err := service.CreateOrder(order)

	if err != nil {
		t.Errorf("unexpected nil, got %v", err)
	}

}
