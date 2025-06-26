package kafka

import (
	"context"
	"encoding/json"

	"github.com/LootNex/CoffeeService/ClientManager/pkg/models"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
)

type OrderProducer interface {
	Send(Order models.Order) error
}

/*

Данная часть кода использует при тестировании работы метода Send

type KafkaWriter interface {
	WriteMessages(ctx context.Context, msgs ...kafka.Message) error
}

type KafkaProducer struct {
	Writer KafkaWriter
}
*/

type KafkaProducer struct {
	Writer *kafka.Writer
}

func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {

	return &KafkaProducer{
		Writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers:          brokers,
			Topic:            topic,
			Balancer:         &kafka.LeastBytes{},
			RequiredAcks:     int(kafka.RequireAll),
			CompressionCodec: &compress.SnappyCodec,
			BatchSize:        100,
			Async:            false,
		}),
	}
}

func (k KafkaProducer) Send(order models.Order) error {

	msg, err := json.Marshal(order)
	if err != nil {
		return err
	}
	return k.Writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: msg,
		},
	)
}
