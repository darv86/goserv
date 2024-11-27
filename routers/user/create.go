package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/darv86/goserv/internal/database"
	"github.com/darv86/goserv/shared"
)

func Create(apiConfig *shared.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("from user create router")
		var params database.UserCreateParams
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			log.Println(err.Error())
			return
		}
		userDb, err := apiConfig.Queries.UserCreate(r.Context(), params)
		if err != nil {
			log.Println(err.Error())
			return
		}
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(userDb)
	}
}
