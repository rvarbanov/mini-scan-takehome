package main

import (
	"context"
	"log"
	"os"

	"github.com/rvarbanov/mini-scan-takehome/internal/consumer"
	"github.com/rvarbanov/mini-scan-takehome/internal/db"
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

	// TODO: move to env/config
	dbHost := "db"
	dbPort := "5432"
	dbUser := "postgres"
	dbPass := "postgres"
	dbName := "mini_scan"

	db := db.NewDB(
		dbHost,
		dbPort,
		dbUser,
		dbPass,
		dbName,
	)

	proc := processor.New(db)

	cons, err := consumer.New(projectID, subID, proc)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer cons.Close()

	if err := cons.Start(context.Background()); err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}
}
