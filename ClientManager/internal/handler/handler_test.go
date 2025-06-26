package handler

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LootNex/CoffeeService/ClientManager/pkg/models"
)

type MockOrderService struct {
	CreateOrderFunc func(order models.Order) error
}

func (m MockOrderService) CreateOrder(order models.Order) error {
	return m.CreateOrderFunc(order)
}

func TestCreateOrderHandler_Success(t *testing.T) {

	mockService := &MockOrderService{
		CreateOrderFunc: func(order models.Order) error {
			return nil
		},
	}

	h := Newhandler(mockService)

	body := `{"id":"1", "buyer":"Y", "item":"latte"}`

	req := httptest.NewRequest(http.MethodPost, "/order", bytes.NewBufferString(body))

	w := httptest.NewRecorder()

	h.CreateOrderHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("cannot read response body: %v", err)
	}

	bodyString := string(bodyBytes)
	t.Logf("response body: %s", bodyString)

}

func TestCreateOrderHandler_BadRequest(t *testing.T) {

	mockService := &MockOrderService{}

	req := httptest.NewRequest(http.MethodPost, "/order", bytes.NewBufferString(`"invalid order"`))

	w := httptest.NewRecorder()

	h := Newhandler(mockService)

	h.CreateOrderHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("cannot read response body err: %v", err)
	}

	bodyString := string(bodyBytes)

	t.Logf("response body: %s", bodyString)

}

func TestCreateOrderHandler_InternalServerError(t *testing.T) {

	mockService := &MockOrderService{
		CreateOrderFunc: func(order models.Order) error {
			return errors.New("ploblems with CreateOrder")
		},
	}

	body := `{"id":"1", "buyer":"Y", "item":"latte"}`

	req := httptest.NewRequest(http.MethodPost, "/order", bytes.NewBufferString(body))
	w := httptest.NewRecorder()

	h := Newhandler(mockService)

	h.CreateOrderHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatalf("cannot read response body err: %v", err)
	}

	bodyString := string(bodyBytes)
	t.Logf("response body: %s", bodyString)

}
