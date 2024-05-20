package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// this will load the env from the .env file
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not in the env")
	}

	// adding a router using chi
	router := chi.NewRouter()

	// set up cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// creating a server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	//start server
	fmt.Println("Listening to server on PORT:", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
