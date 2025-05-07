package config

import (
	"context"
	"os"
)

const (
	EnvKeyHttpPort = "DRONE_HTTP_PORT"
	EnvKeyGrpcPort = "DRONE_GRPC_PORT"

	EnvKeyKafkaHost = "KAFKA_INTERNAL_HOST"
	EnvKeyKafkaPort = "KAFKA_INTERNAL_PORT"
)

type Config struct {
	HttpPort string
	GrpcPort string

	KafkaConfig       *KafkaConfig
	DroneClientConfig *DroneClientConfig
}

func (c *Config) GetKafkaConfig() *KafkaConfig {
	return c.KafkaConfig
}

func (c *Config) GetDroneClientConfig() *DroneClientConfig {
	return c.DroneClientConfig
}

type KafkaConfig struct {
	KafkaHost string
	KafkaPort string
}

func New(ctx context.Context) (*Config, error) {
	return &Config{
		HttpPort: os.Getenv(EnvKeyHttpPort),
		GrpcPort: os.Getenv(EnvKeyGrpcPort),

		KafkaConfig: &KafkaConfig{
			KafkaPort: os.Getenv(EnvKeyKafkaPort),
			KafkaHost: os.Getenv(EnvKeyKafkaHost),
		},

		DroneClientConfig: NewDroneClientConfig(),
	}, nil
}
