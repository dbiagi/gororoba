package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const (
	AppName        = "gororoba"
	DevelopmentEnv = "dev"
	ProductionEnv  = "prod"
)

type Configuration struct {
	WebConfig
	AppConfig
	AWSConfig
}

type AppConfig struct {
	Name        string
	Version     string
	Environment string
}

type WebConfig struct {
	Port            string
	IdleTimeout     time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

type AWSConfig struct {
	Region string
	AWSCredentialsConfig
	DynamoDBConfig
}

type AWSCredentialsConfig struct {
	CredentialType  string
	AccessKeyID     string
	SecretAccessKey string
}

type DynamoDBConfig struct {
	Endpoint string
}

func LoadConfig(env string) Configuration {
	configs := loadFromFile(env)

	return Configuration{
		WebConfig: WebConfig{
			Port:            configs["PORT"],
			IdleTimeout:     time.Second * 10,
			ReadTimeout:     time.Second * 10,
			WriteTimeout:    time.Second * 10,
			ShutdownTimeout: time.Second * 20,
		},
		AppConfig: AppConfig{
			Name:        AppName,
			Version:     "1.0.0",
			Environment: env,
		},
		AWSConfig: AWSConfig{
			Region: configs["AWS_REGION"],
			DynamoDBConfig: DynamoDBConfig{
				Endpoint: configs["AWS_DYNAMODB_ENDPOINT"],
			},
		},
	}
}

func loadFromFile(env string) map[string]string {
	path, _ := os.Getwd()
	configFilePath := fmt.Sprintf("%s/.%s.env", path, env)

	configs, err := godotenv.Read(configFilePath)

	if err != nil {
		slog.Error("Error loading .env file.", slog.String("error", err.Error()))
		panic(err)
	}

	return configs
}
