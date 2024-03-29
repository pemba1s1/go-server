package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pemba1s1/go-server/handlers"
	"github.com/pemba1s1/go-server/internal/database"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port not found in env")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not found in env")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cannot Connect To Database")
	}

	apiCfg := handlers.ApiConfig{
		DB: database.New(conn),
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

	v1Router := chi.NewRouter()
	v1Router.Get("/health", handlers.HandlerReadiness)
	v1Router.Get("/error", handlers.HandlerError)
	v1Router.Post("/user", apiCfg.HandlerCreateUser)
	v1Router.Post("/login", apiCfg.HandlerUserLogin)

	// v1Router.Get("/protected", middleware.AuthMiddleware())

	router.Mount("/v1", v1Router)

	srv := http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	fmt.Printf("Server Starting On Port %v", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server Starting On Port %v", port)
}
