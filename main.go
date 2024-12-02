package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/darv86/goserv/internal/database"
	"github.com/darv86/goserv/routers"
	"github.com/darv86/goserv/scraper"
	"github.com/darv86/goserv/shared"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

func indexRouter(w http.ResponseWriter, r *http.Request) {
	log.Println("from index router")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hi"))
}

func main() {
	connection, err := sql.Open(dbConfig.driver, dbConfig.GetConfigString())
	if err != nil {
		log.Println(err.Error())
	}
	queries := database.New(connection)
	//
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	router.Get("/", indexRouter)
	apiConf := &shared.ApiConfig{Queries: queries, Router: router}
	routers.Setup(apiConf)
	//
	const TICK_INTERVAL = time.Second * 5
	const MAX_FEEDS_AT_TIME = 10
	scraperConf := &shared.ScraperConfig{
		Queries:        queries,
		TickInterval:   TICK_INTERVAL,
		MaxFeedsAtTime: MAX_FEEDS_AT_TIME,
	}
	go scraper.Run(scraperConf)
	//
	PORT := "8080"
	log.Printf("port: %s", PORT)
	http.ListenAndServe(":"+PORT, router)
}
