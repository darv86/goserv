package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainRouter(w http.ResponseWriter, r *http.Request) {
	log.Println("from router main:", r.Host)
	fmt.Fprint(w, "from router main:", r.Host)
}

func apiRouter(w http.ResponseWriter, r *http.Request) {
	log.Println("from router api: ", r.Host)
	fmt.Fprint(w, "from router api: ", r.Host)
}

func main() {
	PORT := "8080"
	http.HandleFunc("/", mainRouter)
	http.HandleFunc("/api", apiRouter)
	http.ListenAndServe(":"+PORT, nil)
	log.Printf("port: %s", PORT)
}
