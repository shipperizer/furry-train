package worker

import (
	"context"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"

	"github.com/shipperizer/furry-train/types"
)

type Consumer struct {
	groupID string
	reader  *kafka.Reader
}

func (consumer *Consumer) Start() {
	log.Info("Listening to kafka on topic: ", consumer.reader.Stats().Topic)

	for {
		msg, err := consumer.reader.ReadMessage(context.Background())

		if err != nil {
			log.Error("Error reading message:", err)
			continue
		}

		event := types.Event{}
		err = proto.Unmarshal(msg.Value, &event)

		if err != nil {
			log.Error("Broken message:", err)
			continue
		}

		log.Info(event.String())
	}

	consumer.reader.Close()
}

func (consumer *Consumer) Close() {
	consumer.reader.Close()
}

func NewConsumer(config ReaderConfig) ConsumerInterface {
	return &Consumer{
		groupID: config.GroupID,
		reader:  NewReader(config),
	}
}
