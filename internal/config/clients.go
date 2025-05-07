package config

import "os"

const (
	EnvKeyDroneClientHost     = "DRONE_CLIENT_HOST"
	EnvKeyDroneClientGrpcPort = "DRONE_CLIENT_GRPC_PORT"
)

type DroneClientConfig struct {
	Host     string
	GrpcPort string
}

func NewDroneClientConfig() *DroneClientConfig {
	return &DroneClientConfig{
		Host:     os.Getenv(EnvKeyDroneClientHost),
		GrpcPort: os.Getenv(EnvKeyDroneClientGrpcPort),
	}
}
