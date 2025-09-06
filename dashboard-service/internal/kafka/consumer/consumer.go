package consumer

import (
	"context"
	config "dashboard-service/internal/config"
	"dashboard-service/internal/repository"
	"dashboard-service/internal/repository/models"
	"dashboard-service/logger"
	"encoding/json"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

type IConsumeInit interface {
	ConsumeMessage() error
	Close() error
}

type ConsumeInit struct {
	client     *kgo.Client
	topic      string
	reportRepo repository.IDashboardRepository
}

func NewConsumeInit(cfg *config.Config, repo repository.IDashboardRepository) (IConsumeInit, error) {

	address := fmt.Sprintf("%s:%d", cfg.Kafka.Host, cfg.Kafka.Port)
	client, err := kgo.NewClient(
		kgo.SeedBrokers(address),
		kgo.ConsumeTopics(cfg.Kafka.Topic),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtEnd()), // Eng oxirgi offsetdan boshlash
	)
	if err != nil {
		logger.Error("Failed to create Kafka client", "error", err)
		return nil, err
	}
	logger.Info("Kafka consumer initialized successfully")
	return &ConsumeInit{
		client:     client,
		topic:      cfg.Kafka.Topic,
		reportRepo: repo,
	}, nil
}

func (c *ConsumeInit) ConsumeMessage() error {

	ctx := context.Background()

	for {
		fetches := c.client.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			for _, err := range errs {
				logger.Error("Error consuming messages", "error", err)
			}
			return fmt.Errorf("error consuming messages: %v", errs)
		}

		fetches.EachPartition(func(partition kgo.FetchTopicPartition) {
			for _, record := range partition.Records {
				logger.Info(fmt.Sprintf("Received message - Key: %s, Value: %s", string(record.Key), string(record.Value)))

				productSales := models.ProductSalesUpdateRequest{}
				if err := json.Unmarshal(record.Value, &productSales); err != nil {
					logger.Error("Error unmarshalling message", "error", err, "message", string(record.Value))
					return
				}

				if err := c.reportRepo.UpsertProductSales(productSales); err != nil {
					logger.Error("Error upserting product sales", "error", err, "product_id", productSales.ProductId)
					return
				}
			}
		})
	}
}

func (c *ConsumeInit) Close() error {
	c.client.Close()
	logger.Info("Kafka client closed successfully")
	return nil
}
