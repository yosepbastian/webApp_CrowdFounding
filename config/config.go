package config

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	UseCache       bool
	TemplateCached map[string]*template.Template
	InfoLog        *log.Logger
	DataSourceName string
}

func (c *Config) ReadConfig() {

	err := godotenv.Load()

	if err != nil {
		panic(err.Error())
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	c.DataSourceName = dsn

}

func NewConfig() Config {
	config := Config{}
	config.ReadConfig()
	return config
}
