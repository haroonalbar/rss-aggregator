package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

	db := database.New(conn)

	// create a api config
	apiCfg := apiConfig{
		DB: db,
	}

  // start scrapping in new go routine 
	go startScrapping(db, 10, time.Minute)

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

	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/error", handlerError)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Put("/feeds/{feedID}", apiCfg.middlewareAuth(apiCfg.handlerUpdateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollows))
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))
  v1Router.Get("/posts",apiCfg.middlewareAuth(apiCfg.handlerGetPostForUser))

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
