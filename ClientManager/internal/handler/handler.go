package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/LootNex/CoffeeService/ClientManager/internal/models"
	"github.com/LootNex/CoffeeService/ClientManager/internal/service"
)

type Handler struct {
	OrderService service.OrderCreator
}

func Newhandler(s service.OrderCreator) Handler {
	return Handler{
		OrderService: s,
	}
}

func (h Handler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {

	order := new(models.Order)

	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		http.Error(w, "cannot decode order", http.StatusBadRequest)
		return
	}

	err = h.OrderService.CreateOrder(*order)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot send order err:%v", err), http.StatusInternalServerError)
		return
	}

	log.Println("order created")

}
