package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	// This fails because PORT is not in the current shell env.
	// It's not taken from the .env file
	// we can add the env by using export command
	// $ export PORT=8000
	// the above command will add the env variable to the current shell
	// but to get it from the .env file we are going to use the godotenv package.

	// this will load the env from the .env file
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not in the env")
	}

	// adding a router
	router := chi.NewRouter()

	// creating a server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Println("Listening to server on PORT:", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening to server on PORT:", port)
}
