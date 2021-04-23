package conf

import (
	"github.com/joho/godotenv"

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
	// log.Info("Environment: ", os.Getenv("ENVIRONMENT"))
	if os.Getenv("ENVIRONMENT") == EnvironmentLocal {
		err := godotenv.Load()
		if err != nil {
			// log.Error("Error loading .env file")
			return err
		}
	}

	// log.Info("RunMode: ", EnvConfig.RunMode)

	EnvConfig.MySQL.Host = os.Getenv("MYSQL_HOST")
	mysqlPort, err := strconv.ParseInt(os.Getenv("MYSQL_PORT"), 10, 64)
	if err != nil {
		// log.Error("Error when parse config MYSQL_PORT, detail: ", err)
		return err
	}
	EnvConfig.MySQL.Port = mysqlPort
	EnvConfig.MySQL.User = os.Getenv("MYSQL_USER")
	EnvConfig.MySQL.Password = os.Getenv("MYSQL_PASSWORD")
	EnvConfig.MySQL.DB = os.Getenv("MYSQL_DBNAME")

	EnvConfig.Server.BindingHost = os.Getenv("BINDING_HOST")
	serverBindingPort, err := strconv.ParseInt(os.Getenv("BINDING_PORT"), 10, 64)
	if err != nil {
		// log.Error("Error when parse config BINDING_PORT, detail: ", err)
		return err
	}
	EnvConfig.Server.BindingPort = serverBindingPort

	EnvConfig.Environment = os.Getenv("ENVIRONMENT")

	EnvConfig.JWT.PublicKey = os.Getenv("JWT_PUBLIC_KEY")

	return nil
}
