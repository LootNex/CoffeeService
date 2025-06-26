package kafka

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/segmentio/kafka-go"
)

type MockReadKafkaReader struct {
	Messages []kafka.Message
}

func (m *MockReadKafkaReader) Close() error {
	return nil
}

func (m *MockReadKafkaReader) ReadMessage(ctx context.Context) (kafka.Message, error) {

	if len(m.Messages) == 0 {
		return kafka.Message{}, errors.New("slice messages is empty")
	}

	msg := m.Messages[0]
	m.Messages = m.Messages[1:]

	return msg, nil

}

func TestKafkaConsumer(t *testing.T) {

	mock := &MockReadKafkaReader{
		Messages: []kafka.Message{
			{Value: []byte("latte")},
			{Value: []byte("americano")},
		},
	}

	var received []string

	handle := func(msg kafka.Message) {
		received = append(received, string(msg.Value))
	}

	err := KafkaConsumer(mock, handle)

	if err.Error() != "slice messages is empty" && err != nil {
		t.Errorf("unexpected err %v", err)
	}

	expected := []string{"latte", "americano"}

	if !reflect.DeepEqual(received, expected) {
		t.Errorf("expected %v, got %v", expected, received)
	}

}
