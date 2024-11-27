package user

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	"github.com/darv86/goserv/internal/utils"
	"github.com/darv86/goserv/shared"
)

func GetById(apiConf *shared.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("from get user by id router")
		id, err := utils.GetUrlLastParam(r.URL)
		if err != nil {
			log.Println(err.Error())
			return
		}
		w.Header().Add("Content-type", "application/json")
		user := apiConf.AuthUser
		if !reflect.ValueOf(user).IsZero() && user.ID == id {
			json.NewEncoder(w).Encode(user)
			return
		}
		userDb, err := apiConf.Queries.UserGetById(r.Context(), id)
		if err != nil {
			log.Println(err.Error())
		}
		json.NewEncoder(w).Encode(userDb)
	}
}
