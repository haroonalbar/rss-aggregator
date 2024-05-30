package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/haroonalbar/rss-aggregater/internal/database"
	"github.com/joho/godotenv"

	// have to do this to for the db connection to work properly
	// this is a db driver
	// but we don't have to user it
	_ "github.com/lib/pq"
)

// database Queries is taken from the generated sqlc code.
type apiConfig struct {
	DB *database.Queries
}

func main() {

	// this will load the env from the .env file
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not in the env")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not in the env")
	}

	//connect to db
	// postgres is the driver name
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Unalbe to connect to db:", err)
	}

	// create a api config
	apiCfg := apiConfig{
		DB: database.New(conn),
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
	v1Router.Get("/ready", handlerReadiness)
	//handle error
	v1Router.Get("/error", handlerError)
	//handle create user
	v1Router.Post("/users", apiCfg.handlerCreateUser)

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
