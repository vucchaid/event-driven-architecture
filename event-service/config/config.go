package config

import (
	"log"
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
	log.Println("Loaded PORT=", config.Port)
	config.DBtype = os.Getenv("DBTYPE")
	log.Println("Loaded DBTYPE=", config.DBtype)
	config.Connection = os.Getenv("CONN_URL")
	log.Println("Loaded CONN_URL=", config.Connection)

	return &config, nil
}

func LoadEnvironmentVariables() error {
	return godotenv.Load(".env")
}
