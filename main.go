package main

import (
	"database/sql"
	"log"
	"net/http"

	// to use external libs:
	// 1. go get github.com/go-chi/chi/v5 (package path)
	// 2. import and use some code in this module
	// 3. go mod vendor (folder for a libs)
	// 4. go mod tidy (to clean/fix all requires in go.mod)
	"github.com/darv86/goserv/internal/database"
	"github.com/darv86/goserv/routers"
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
	routers.Setup(router, queries)
	//
	PORT := "8080"
	log.Printf("port: %s", PORT)
	http.ListenAndServe(":"+PORT, router)
}
