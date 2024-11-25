package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/darv86/goserv/internal/database"
	"github.com/go-chi/chi/v5"
)

func DeleteById(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("from delete user by id router")
		param := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			log.Println(err.Error())
			return
		}
		userDb, err := queries.UserDeleteById(r.Context(), id)
		if err != nil {
			log.Println(err.Error())
		}
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(userDb)
	}
}
