package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/darv86/goserv/shared"
)

func GetAll(apiConf *shared.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("from get users router")
		usersDb, err := apiConf.Queries.UserGetAll(r.Context())
		if err != nil {
			log.Println(err.Error())
		}
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(usersDb)
	}
}
