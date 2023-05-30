package main

import (
	"fmt"
	"log"
	"net/http"
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
	// r.Get("/stats", Stats)

	// --- LOGS ---
	log.Printf("Server strating on port %d...\n", config.HttpPort)
	log.Printf("Env: %s\n", config.Env)

	// --- "DB" ---
	// Creating the simplest "DB": a csv file as I don't have time to implement a Dockerized DB and using squirrel.
	data.InitialisingPseudoDB(config.DBFilename)

	// Launching server
	if err := http.ListenAndServe(s.Addr, s.Handler); err != nil {
		log.Fatal(fmt.Errorf("server seemingly crashed, more on that: %w", err))
	}
}
