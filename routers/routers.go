package routers

import (
	"errors"
	"log"
	"net/http"

	"github.com/darv86/goserv/internal/utils"
	"github.com/darv86/goserv/routers/feed"
	"github.com/darv86/goserv/routers/user"
	"github.com/darv86/goserv/shared"
)

type AuthMiddlewareHandler func(*shared.ApiConfig) http.HandlerFunc

func AuthMiddleware(apiConf *shared.ApiConfig, handler AuthMiddlewareHandler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		unAuthErr := errors.New("not authorized")
		apiKey, err := utils.GetApiKey(req.Header)
		if err != nil {
			log.Println(err.Error())
			http.Error(res, unAuthErr.Error(), http.StatusUnauthorized)
			return
		}
		userDb, err := apiConf.Queries.UserGetByApiKey(req.Context(), apiKey)
		if err != nil {
			log.Println(err.Error())
			http.Error(res, unAuthErr.Error(), http.StatusUnauthorized)
			return
		}
		apiConf.AuthUser = userDb
		hf := handler(apiConf)
		hf(res, req)
	}
}

func Setup(apiConf *shared.ApiConfig) {
	router := apiConf.Router
	router.Get("/users", user.GetAll(apiConf))
	router.Get("/user/{id}", AuthMiddleware(apiConf, user.GetById))
	router.Post("/user/create", user.Create(apiConf))
	router.Delete("/users/delete", AuthMiddleware(apiConf, user.DeleteAll))
	router.Delete("/user/delete/{id}", AuthMiddleware(apiConf, user.DeleteById))

	router.Get("/feeds", feed.GetAll(apiConf))
	router.Post("/feed/create", AuthMiddleware(apiConf, feed.Create))
	router.Delete("/feeds/delete", AuthMiddleware(apiConf, feed.DeleteAll))

	router.Get("/feeds-mine", AuthMiddleware(apiConf, feed.MineGetAll))
	router.Delete("/feed-mine/delete/{id}", AuthMiddleware(apiConf, feed.MineDeleteById))
}
