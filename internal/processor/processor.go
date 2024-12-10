package processor

import (
	"fmt"

	"cloud.google.com/go/pubsub"
)

type Processor struct{}

func New() *Processor {
	return &Processor{}
}

func (p *Processor) ProcessMessage(msg *pubsub.Message) error {
	fmt.Printf("Processing message: %s\n", string(msg.Data))
	// TODO:
	// store message data into the db

	return nil
}
