package main

import (
	"context"

	"github.com/kelseyhightower/envconfig"
	uuid "github.com/satori/go.uuid"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/shipperizer/furry-train/types"
)

type EnvSpec struct {
	Topic   string   `envconfig:"topic" default:"protobuf"`
	Brokers []string `envconfig:"brokers"`
}

func NewProducer(brokers []string, topic string) *kafka.Writer {
	return kafka.NewWriter(
		kafka.WriterConfig{
			Brokers:  brokers,
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
			Async:    true,
			Dialer: &kafka.Dialer{
				DualStack: true,
			},
		},
	)
}

func Spam(producer *kafka.Writer) {
	events := make([]kafka.Message, 0)
	for i := 0; i < 100; i++ {
		UUID := uuid.NewV4().String()
		cfg, _ := structpb.NewStruct(
			map[string]interface{}{
				"id":   int64(i),
				"uuid": UUID,
			},
		)

		event := types.Event{
			Uuid:   UUID,
			Id:     int64(i),
			Config: cfg,
		}

		msg, _ := proto.Marshal(&event)

		events = append(
			events,
			kafka.Message{
				Key:   []byte(event.GetUuid()),
				Value: msg,
			},
		)
	}

	producer.WriteMessages(context.TODO(), events...)
	producer.Close()
}

func main() {
	var specs EnvSpec
	err := envconfig.Process("", &specs)

	if err != nil {
		log.Fatal(err.Error())
	}

	producer := NewProducer(specs.Brokers, specs.Topic)

	Spam(producer)
}
