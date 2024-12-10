package env

import (
	"fmt"
	"os"
)

type Env struct {
	PubSub struct {
		ProjectID string
		SubID     string
	}
	DB struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}
}

type PubSubEnv struct {
	ProjectID string
	SubID     string
}

type DBEnv struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

func GetEnvs() Env {
	projectID := os.Getenv("PUBSUB_PROJECT_ID")
	validateEnv("PUBSUB_PROJECT_ID", projectID)
	subID := os.Getenv("PUBSUB_SUB_ID")
	validateEnv("PUBSUB_SUB_ID", subID)

	dbHost := os.Getenv("POSTGRES_HOST")
	validateEnv("POSTGRES_HOST", dbHost)
	dbPort := os.Getenv("POSTGRES_PORT")
	validateEnv("POSTGRES_PORT", dbPort)
	dbUser := os.Getenv("POSTGRES_USER")
	validateEnv("POSTGRES_USER", dbUser)
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	validateEnv("POSTGRES_PASSWORD", dbPass)
	dbName := os.Getenv("POSTGRES_DB_NAME")
	validateEnv("POSTGRES_DB_NAME", dbName)

	return Env{
		PubSub: PubSubEnv{
			ProjectID: projectID,
			SubID:     subID,
		},
		DB: DBEnv{
			Host: dbHost,
			Port: dbPort,
			User: dbUser,
			Pass: dbPass,
			Name: dbName,
		},
	}
}

func validateEnv(key string, value string) {
	if value == "" {
		panic(fmt.Sprintf("%s is not set", key))
	}
}
