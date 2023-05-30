package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DBFilename string
	Env        string
	HttpPort   int
)

func Init() {
	//Loading data from .env file or data (AWS and such...)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Env
	Env = os.Getenv("ENV")
	if Env == "" {
		Env = "development"
	}

	// HttpPort
	if HttpPort, err = strconv.Atoi(os.Getenv("SERVER_HTTP_PORT")); err != nil {
		log.Fatal(fmt.Errorf("missing or unreadable server port from env: %w", err))
	}

	if HttpPort == 0 {
		HttpPort = 9000
		log.Printf("setting default server port as: 9000")
	}

	// DB
	DBFilename = os.Getenv("DB_FILENAME")
	if DBFilename == "" {
		DBFilename = "db_fizzbuzz.csv"
	}
}
