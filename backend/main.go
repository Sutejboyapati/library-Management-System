package main

import (
	"log"
	"net/http"
	"os"

	"library/backend/config"
	"library/backend/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := config.InitDB(); err != nil {
		log.Fatalf("Database init failed: %v", err)
	}

	router := mux.NewRouter()
	routes.Setup(router)
	r := http.NewServeMux()
	r.Handle("/", router)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(r)

	log.Printf("Server listening on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
