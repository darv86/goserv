package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/darv86/goserv/internal/database"
)

func CreateUserRouter(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("from user create router")
		//
		var params database.CreateUserParams
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			log.Println(err.Error())
			return
		}
		//
		user, err := queries.CreateUser(r.Context(), params)
		if err != nil {
			log.Println(err.Error())
			return
		}
		//
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}
