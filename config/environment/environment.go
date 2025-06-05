package environment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

const (
	ConnectionString string = "CONNECTION_STRING"
	AutoMigrate      string = "AUTO_MIGRATE"
	Port             string = "PORT"
)

type DataBaseConfig struct {
	ConnectionString string
	AutoMigrate      bool
}

type ServerConfig struct {
	Port string
}

func NewDatabaseConfig() *DataBaseConfig {
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Sprintf("env file is not well configured: %s", err.Error()))
	}

	return &DataBaseConfig{
		ConnectionString: os.Getenv(ConnectionString),
		AutoMigrate:      os.Getenv(AutoMigrate) == "true",
	}
}

func NewServerConfig() *ServerConfig {
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Sprintf("env file is not well configured: %s", err.Error()))
	}

	return &ServerConfig{
		Port: os.Getenv(Port),
	}
}

func ProvideEnvironment() fx.Option {
	return fx.Module(
		"env",
		fx.Provide(
			NewDatabaseConfig,
			NewServerConfig,
		),
	)
}
