package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type Page struct {
	Title string
	Body  string
}

func loadPage(name string) (*Page, error) {
	filePath := "db/" + name + ".txt"
	body, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return &Page{Title: name, Body: string(body)}, nil
}

func renderTemplate(w http.ResponseWriter, name string, page *Page) {
	t, _ := template.ParseFiles("static/" + name + ".html")
	t.Execute(w, page)
}

func indexRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")[1]
	if len(path) == 0 {
		path = "index"
	}
	log.Println("path from index:", path)
	page, err := loadPage(path)
	if err != nil {
		page = &Page{Title: "Wrong page title content", Body: "Wrong page body content"}
	}
	renderTemplate(w, "page", page)
	log.Println("router index:", r.Host, r.URL.Path)
}

func loginRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")[1]
	if len(path) == 0 {
		path = "index"
	}
	log.Println("path from login:", path)
	page, err := loadPage(path)
	if err != nil {
		page = &Page{Title: "Wrong page title content", Body: "Wrong page body content"}
	}
	renderTemplate(w, "page", page)
	log.Println("router login:", r.Host, r.URL.Path)
}

func main() {
	PORT := "8080"
	http.HandleFunc("/", indexRouter)
	http.HandleFunc("/login", loginRouter)
	http.ListenAndServe(":"+PORT, nil)
	log.Printf("port: %s", PORT)
}
