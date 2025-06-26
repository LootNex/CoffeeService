package kafka

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/LootNex/CoffeeService/ClientManager/pkg/models"
	"github.com/segmentio/kafka-go"
)

type MockWriter struct {
	WriteMessageFunc func(ctx context.Context, msgs ...kafka.Message) error
}

func (m MockWriter) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	return m.WriteMessageFunc(ctx, msgs...)
}

func TestSend(t *testing.T) {

	var sentMessage []byte

	// mock := MockWriter{
	// 	WriteMessageFunc: func(ctx context.Context, msgs ...kafka.Message) error {
	// 		if len(msgs) != 1 {
	// 			t.Errorf("expected 1 message, got %d", len(msgs))
	// 		}
	// 		sentMessage = msgs[0].Value
	// 		return nil
	// 	},
	// }

	producer := KafkaProducer{
		// Writer: mock,
	}

	order := models.Order{
		ID:    "1",
		Buyer: "N",
		Item:  "latte",
	}

	err := producer.Send(order)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedJSON, _ := json.Marshal(order)
	if string(sentMessage) != string(expectedJSON) {
		t.Errorf("expected message %s, got %s", expectedJSON, sentMessage)
	}
}
