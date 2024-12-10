package main

import (
	"context"
	"log"
	"os"

	"github.com/rvarbanov/mini-scan-takehome/internal/consumer"
	"github.com/rvarbanov/mini-scan-takehome/internal/processor"
)

// TODO: move to env/config
const (
	pubSubEmulatorHost = "pubsub:8085"
	projectID          = "test-project"
	subID              = "scan-sub"
)

func main() {
	// Set pubsub emulator host
	os.Setenv("PUBSUB_EMULATOR_HOST", pubSubEmulatorHost)

	proc := processor.New()

	cons, err := consumer.New(projectID, subID, proc)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer cons.Close()

	if err := cons.Start(context.Background()); err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}
}
