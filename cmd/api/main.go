package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/pemba1s1/go-server/handlers"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port not found in env")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Rounter := chi.NewRouter()
	v1Rounter.Get("/health", handlers.HandlerReadiness)

	router.Mount("/v1", v1Rounter)

	srv := http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	fmt.Printf("Server Starting On Port %v", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server Starting On Port %v", port)
}
