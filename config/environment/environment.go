package environment

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

const (
	ConnectionString string = "CONNECTION_STRING"
	AutoMigrate      string = "AUTO_MIGRATE"
	Port             string = "PORT"
	Host             string = "HOST"
	SecretKey        string = "PASSWORD_SECRET_KEY"
	AutoSeed         string = "AUTO_SEED"
)

type DataBaseConfig struct {
	ConnectionString string
	AutoMigrate      bool
}

type ServerConfig struct {
	Host string
	Port string
}

type JwtConfig struct {
	Secret          string
	Issuer          string
	Audience        string
	ExpirationHours int
}

type PasswordConfig struct {
	SecretKey string
}

type AdminUserConfig struct {
	Email    string
	Password string
}

type SeedConfig struct {
	AutoSeed bool
}

func NewSeedConfig() *SeedConfig {
	return &SeedConfig{
		AutoSeed: os.Getenv(AutoSeed) == "true",
	}
}

func NewPasswordConfig() *PasswordConfig {
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Sprintf("env file is not well configured: %s", err.Error()))
	}

	return &PasswordConfig{
		SecretKey: os.Getenv("PASSWORD_SECRET_KEY"),
	}
}

func NewAdminUserConfig() *AdminUserConfig {
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Sprintf("env file is not well configured: %s", err.Error()))
	}

	return &AdminUserConfig{
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: os.Getenv("ADMIN_PASSWORD"),
	}
}

func NewJwtConfig() *JwtConfig {
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Sprintf("env file is not well configured: %s", err.Error()))
	}

	hours, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_HOURS"))
	if err != nil {
		panic(fmt.Sprintf("env file is not well configured: %s", err.Error()))
	}
	return &JwtConfig{
		Secret:          os.Getenv("JWT_SECRET"),
		Issuer:          os.Getenv("JWT_ISSUER"),
		Audience:        os.Getenv("JWT_AUDIENCE"),
		ExpirationHours: hours,
	}
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
		Host: os.Getenv(Host),
	}
}

func ProvideEnvironment() fx.Option {
	return fx.Module(
		"env",
		fx.Provide(
			NewDatabaseConfig,
			NewServerConfig,
			NewPasswordConfig,
			NewSeedConfig,
			NewJwtConfig,
		),
	)
}
