package conf

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"os"
	"strconv"
)

const (
	EnvironmentLocal = "LOCAL"
)

// Config db environment
type Config struct {
	MySQL struct {
		Host     string
		Port     int64
		User     string
		Password string
		DB       string
	}

	Server struct {
		BindingHost string
		BindingPort int64
	}

	JWT struct {
		PrivateKey string
		PublicKey  string
	}

	Environment string
}

// EnvConfig env config
var EnvConfig *Config

// init function
func InitConfig() error {
	EnvConfig = &Config{}
	EnvConfig.Environment = os.Getenv("ENVIRONMENT")
	zap.S().Info("Environment: ", EnvConfig.Environment)
	if EnvConfig.Environment == EnvironmentLocal {
		err := godotenv.Load()
		if err != nil {
			zap.S().Error("Error loading .env file")
			return err
		}
	}

	EnvConfig.MySQL.Host = os.Getenv("MYSQL_HOST")
	mysqlPort, err := strconv.ParseInt(os.Getenv("MYSQL_PORT"), 10, 64)
	if err != nil {
		zap.S().Error("Error when parse config MYSQL_PORT, detail: ", err)
		return err
	}
	EnvConfig.MySQL.Port = mysqlPort
	EnvConfig.MySQL.User = os.Getenv("MYSQL_USER")
	EnvConfig.MySQL.Password = os.Getenv("MYSQL_PASSWORD")
	EnvConfig.MySQL.DB = os.Getenv("MYSQL_DBNAME")

	EnvConfig.Server.BindingHost = os.Getenv("BINDING_HOST")
	serverBindingPort, err := strconv.ParseInt(os.Getenv("BINDING_PORT"), 10, 64)
	if err != nil {
		zap.S().Error("Error when parse config BINDING_PORT, detail: ", err)
		return err
	}
	EnvConfig.Server.BindingPort = serverBindingPort

	EnvConfig.JWT.PublicKey = os.Getenv("JWT_PUBLIC_KEY")
	EnvConfig.JWT.PrivateKey = os.Getenv("JWT_PRIVATE_KEY")

	return nil
}
