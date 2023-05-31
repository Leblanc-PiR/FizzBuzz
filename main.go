package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/Leblanc-PiR/FizzBuzz/config"
	"github.com/Leblanc-PiR/FizzBuzz/internal/data"
	"github.com/Leblanc-PiR/FizzBuzz/internal/handler"
)

func main() {
	// All of this could run off of a Docker or Kubernetes

	config.Init()

	// Chi: never used it, trying.

	// --- ROUTER ---
	// Init Router
	r := chi.NewRouter()

	/// Setting CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET"}, // reducing usage for security purpose
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           200,
	})
	r.Use(cors.Handler)

	// Server config
	s := &http.Server{
		Addr:           ":" + strconv.Itoa(config.HttpPort),
		Handler:        chi.NewMux().Middlewares().Handler(r),
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Logger
	r.Use(middleware.Logger)

	// --- ROUTES ---
	// Health
	r.Get("/health", handler.Health)

	// FizzBuzz
	r.Get("/fizzbuzz", handler.GetFizzBuzz)

	// Stats (could be secured behind some BasicAuth at least)
	r.Get("/stats", handler.GetStats)

	// --- LOGS ---
	log.Printf("Server starting on port %d...\n", config.HttpPort)
	log.Printf("Env: %s\n", config.Env)

	// --- "DB" ---
	// Creating the simplest "DB": a json file.
	// As implementing a Dockerized DB and using squirrel would have been time consuming.
	data.InitialisingDB(config.DBFilename)

	// ServerErrors channel
	serverErrors := make(chan error, 1)

	// Launching server
	go func() {
		serverErrors <- http.ListenAndServe(s.Addr, s.Handler)
	}()

	// Softer crash
	err := <-serverErrors
	if err != nil {
		log.Panic(fmt.Errorf("server seemingly crashed, more on that: %w", err))
		os.Exit(1)
	}

}
