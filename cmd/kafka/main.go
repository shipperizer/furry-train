package main

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"

	"github.com/shipperizer/furry-train/worker"
)

type EnvSpec struct {
	Topic   string   `envconfig:"topic" default:"protobuf"`
	Brokers []string `envconfig:"brokers"`
	GroupID string   `envconfig:"cgroup" default:"protobuf"`
}

func main() {
	var specs EnvSpec
	err := envconfig.Process("", &specs)

	if err != nil {
		log.Fatal(err.Error())
	}

	cfg := worker.ReaderConfig{
		Topic:   specs.Topic,
		Brokers: specs.Brokers,
		GroupID: specs.GroupID,
	}

	consumer := worker.NewConsumer(cfg)
	consumer.Start()
}
