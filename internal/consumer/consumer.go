package consumer

import (
	"context"
	"fmt"
	"log"

	"github.com/rvarbanov/mini-scan-takehome/internal/processor"

	"cloud.google.com/go/pubsub"
)

type Consumer struct {
	client    *pubsub.Client
	sub       *pubsub.Subscription
	processor *processor.Processor
}

func New(projectID string, subID string, processor *processor.Processor) (*Consumer, error) {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create pubsub client: %w", err)
	}

	sub := client.Subscription(subID)

	return &Consumer{
		client:    client,
		sub:       sub,
		processor: processor,
	}, nil
}

func (c *Consumer) Start(ctx context.Context) error {
	return c.sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		if err := c.processor.ProcessMessage(msg); err != nil {
			log.Printf("Failed to process message: %v", err)
			msg.Nack()
		} else {
			msg.Ack()
		}
	})
}

func (c *Consumer) Close() {
	c.client.Close()
}
