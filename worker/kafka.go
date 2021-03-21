package worker

import (
	"github.com/segmentio/kafka-go"
)

type ReaderConfig struct {
	Brokers []string
	Topic   string
	GroupID string
}

func NewReader(cfg ReaderConfig) *kafka.Reader {
	return kafka.NewReader(
		kafka.ReaderConfig{
			Brokers:  cfg.Brokers,
			GroupID:  cfg.GroupID,
			Topic:    cfg.Topic,
			MinBytes: 10e3, // 10KB
			MaxBytes: 10e6, // 10MB,
			Dialer: &kafka.Dialer{
				DualStack: true,
			},
		},
	)
}
