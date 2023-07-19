package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ApplicationConfig struct {
	Port, Connection, DBtype string
}

func GetApplicationConfig() (*ApplicationConfig, error) {

	err := LoadEnvironmentVariables()
	if err != nil {
		return nil, err
	}

	config := ApplicationConfig{}

	config.Port = os.Getenv("PORT")
	config.DBtype = os.Getenv("DBTYPE")
	config.Connection = os.Getenv("CONN_URL")

	return &config, nil
}

func LoadEnvironmentVariables() error {
	return godotenv.Load("somerandomfile")
}
