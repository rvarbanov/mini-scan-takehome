package main

import (
	"context"
	"log"

	"github.com/rvarbanov/mini-scan-takehome/internal/consumer"
	"github.com/rvarbanov/mini-scan-takehome/internal/db"
	"github.com/rvarbanov/mini-scan-takehome/internal/env"
	"github.com/rvarbanov/mini-scan-takehome/internal/processor"
)

// TODO: move to env/config
const (
/*
projectID = "test-project"
subID     = "scan-sub"

dbHost = "db"
dbPort = "5432"
dbUser = "postgres"
dbPass = "postgres"
dbName = "mini_scan"
*/
)

func main() {
	envs := env.GetEnvs()

	db := db.NewDB(
		envs.DB.Host,
		envs.DB.Port,
		envs.DB.User,
		envs.DB.Pass,
		envs.DB.Name,
	)

	proc := processor.New(db)

	cons, err := consumer.New(envs.PubSub.ProjectID, envs.PubSub.SubID, proc)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer cons.Close()

	if err := cons.Start(context.Background()); err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}
}
