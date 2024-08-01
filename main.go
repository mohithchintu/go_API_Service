package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	Port := os.Getenv("PORT")

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + Port,
	}
	log.Printf("Server running on port %v", Port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
