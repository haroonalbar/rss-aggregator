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

  // create another router to mount on router for versioning
  // this is for future proofing the api
	v1Router := chi.NewRouter()

  // handle the /ready route with heandleReadiness function
  v1Router.Get("/ready",handlerReadiness)
  //handle error 
  v1Router.Get("/error",handlerError)

  // mount the v1Router to router so the whole path will become path/v1/ready
	router.Mount("/v1", v1Router)

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
