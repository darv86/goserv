package routers

import (
	"net/http"

	"github.com/darv86/goserv/internal/database"
	"github.com/darv86/goserv/routers/user"
)

type Router interface {
	Get(pattern string, handlerFn http.HandlerFunc)
	Post(pattern string, handlerFn http.HandlerFunc)
	Delete(pattern string, handlerFn http.HandlerFunc)
}

func Setup(router Router, queries *database.Queries) {
	router.Get("/users", user.GetAll(queries))
	router.Get("/user/{id}", user.GetById(queries))
	router.Post("/user/create", user.Create(queries))
	router.Delete("/users/delete", user.DeleteAll(queries))
	router.Delete("/user/delete/{id}", user.DeleteById(queries))
}
