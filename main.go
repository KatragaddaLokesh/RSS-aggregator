package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/KatragaddaLokesh/RSSagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")
	port := os.Getenv("PORT")

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is Not Found")
	}

	if port == "" {
		log.Fatal("Port is Not Found")
	}

	connec, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Cannot Connect to db", err)
	}

	db := database.New(connec)
	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

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
	v1Router.Get("/healthz", handleReady)
	v1Router.Get("/err", handleError)

	v1Router.Post("/users", apiCfg.handleUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handleGetUser))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handleFeed))
	v1Router.Get("/feeds", apiCfg.handleGetFeed)

	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handleGetPostForUsers))

	v1Router.Post("/feeds_follow", apiCfg.middlewareAuth(apiCfg.handleFeedFollow))
	v1Router.Get("/feeds_follow", apiCfg.middlewareAuth(apiCfg.handleGetFeedFollow))
	v1Router.Delete("/feeds_follow/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handleDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("Server Starting on %v", port)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port", port)
}
