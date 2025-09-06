package producer

import (
	"product-service/logger"
	"product-service/internal/config"
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

type IProducerInit interface {
	ProduceMessage(key string, message []byte) error
	Close() error
}

type ProducerInit struct {
	client *kgo.Client
	topic  string
}


func NewProducerInit(cfg *config.Config) (IProducerInit, error) {

	address := fmt.Sprintf("%s:%d", cfg.Kafka.Host, cfg.Kafka.Port)
	logger.Info("Initializing Kafka producer with address:", address)
	client, err := kgo.NewClient(
		kgo.SeedBrokers(address),
		kgo.AllowAutoTopicCreation(),
	)
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}
	logger.Info("Kafka client initialized successfully")
	return &ProducerInit{
		client: client, 
		topic: cfg.Kafka.Topic,
	}, nil
}

func (p *ProducerInit) ProduceMessage(key string, message []byte) error {

	logger.Info("Producing message to topic:", p.topic, " with key:", key)
	
	record := &kgo.Record{
		Topic: p.topic,
		Key:   []byte(key),
		Value: message,
	}
	
	err := p.client.ProduceSync(context.Background(), record).FirstErr()
	if err != nil {
		logger.Error(err)
		return err
	}
	logger.Info("Message produced successfully to topic:", p.topic)
	return nil
}

func (p *ProducerInit) Close() error {

	logger.Info("Closing Kafka client")
	p.client.Close()
	return nil
}
