package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Defining config vars
var (
	DBFilename string
	Env        string
	HttpPort   int
)

// Init initialises env vars
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
		log.Println("Missing or unreadable server port from .env")

		HttpPort = 9000
		log.Printf("Setting default server port as: %d", HttpPort)
	}

	// "DB"
	DBFilename = os.Getenv("DB_FILENAME") // Could be DB connection, made it lightweight.
	if DBFilename == "" {
		DBFilename = "db_fizzbuzz.json"
	}
}
