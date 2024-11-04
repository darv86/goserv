package main

import (
	"fmt"
	"log"
	"os"

	// to use external libs:
	// 1. go get github.com/go-chi/chi/v5 (package path)
	// 2. import and use some code in this module
	// 3. go mod vendor (folder for a libs)
	// 4. go mod tidy (to clean/fix all requires in go.mod)
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	PORT := "8080"
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	// http.ListenAndServe(":"+PORT, router)
	log.Printf("port: %s", PORT)
	host, _ := os.Getwd()
	fmt.Println("server", host)
}
