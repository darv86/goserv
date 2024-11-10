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

// Template caching
// variable (templates) has all templates which were parsed on a start of this program
// name of each template has the name of the file name (excluding path name);
// if we need a global scope variable ( only for main.go),
// we should use var keyword for this variable,
// but now to use := syntax (only for local scope: e.g. inside a func)
var templates = template.Must(template.ParseFiles("static/page.html"))

func loadPage(name string) (*Page, error) {
	filePath := "db/" + name + ".txt"
	body, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return &Page{Title: name, Body: string(body)}, nil
}

func renderTemplate(w http.ResponseWriter, name string, page *Page) {
	// render specific template by its name
	err := templates.ExecuteTemplate(w, name+".html", page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")[1]
	if len(path) == 0 {
		path = "index"
	}
	page, err := loadPage(path)
	if err != nil {
		// handle an error, using a few options:
		// 1. create Page instance with default info and assign pointer to page variable
		// page = &Page{Title: "Wrong page title content", Body: "Wrong page body content"}
		// 2. Redirect to index.html, if path doesn't matched with any db entities
		// http.Redirect(w, r, "/", http.StatusFound)
		// 3. call error to inform user
		// Error method on an error returns string description of that error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "page", page)
	log.Println("router index:", r.Host, r.URL.Path)
}

func loginRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")[1]
	if len(path) == 0 {
		path = "index"
	}
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
