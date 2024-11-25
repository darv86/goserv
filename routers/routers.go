package routers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/darv86/goserv/internal/database"
	"github.com/darv86/goserv/internal/utils"
	"github.com/darv86/goserv/routers/user"
	"github.com/go-chi/chi/v5"
)

type Router interface {
	Get(pattern string, handlerFn http.HandlerFunc)
	Post(pattern string, handlerFn http.HandlerFunc)
	Delete(pattern string, handlerFn http.HandlerFunc)
}

type Handler func(*database.Queries) http.HandlerFunc

func AuthMiddleware(queries *database.Queries, handler http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		apiKey, err := utils.GetApiKey(req.Header)
		if err != nil {
			log.Println(err.Error())
			http.Error(res, err.Error(), http.StatusUnauthorized)
			return
		}
		userDb, err := queries.UserGetByApiKey(req.Context(), apiKey)
		if err != nil {
			log.Println(err.Error())
			http.Error(res, err.Error(), http.StatusUnauthorized)
			return
		}
		param := chi.URLParam(req, "id")
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			log.Println(err.Error())
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		if userDb.ID != id {
			notAuthErr := errors.New("not authorized request")
			log.Println(notAuthErr)
			http.Error(res, notAuthErr.Error(), http.StatusUnauthorized)
			return
		}
		handler(res, req)
	}
}

func Setup(router Router, queries *database.Queries) {
	router.Get("/users", user.GetAll(queries))
	router.Get("/user/{id}", AuthMiddleware(queries, user.GetById(queries)))
	router.Post("/user/create", user.Create(queries))
	router.Delete("/users/delete", user.DeleteAll(queries))
	router.Delete("/user/delete/{id}", user.DeleteById(queries))
}
