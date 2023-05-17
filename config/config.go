package config

import (
	"log"
	"os"

	//Use this to fetch the package for environment variables
	"github.com/joho/godotenv"
)

// This will act as a file to store all the constants
type Config struct {
	ServerHost string
	MongoURL   string
}

// To set the value of config object
func (c *Config) initialize() {
	//fetch value from local env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't fetch env variables")
		os.Exit(0)
	}

	c.MongoURL = os.Getenv("MONGO_URL")
	c.ServerHost = os.Getenv("PORT")

}

func (c *Config) GetMongoURL() string {
	return c.MongoURL
}

func NewConfig() *Config {
	config := new(Config)
	config.initialize()
	return config
}
