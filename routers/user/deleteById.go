package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/darv86/goserv/internal/utils"
	"github.com/darv86/goserv/shared"
)

func DeleteById(apiConf *shared.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("from delete user by id router")
		id, err := utils.GetUrlLastParam(r.URL)
		if err != nil {
			log.Println(err.Error())
			return
		}
		userDb, err := apiConf.Queries.UserDeleteById(r.Context(), id)
		if err != nil {
			log.Println(err.Error())
		}
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(userDb)
	}
}
