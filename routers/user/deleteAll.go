package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/darv86/goserv/shared"
)

func DeleteAll(apiConf *shared.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("from delete users router")
		err := apiConf.Queries.UserDeleteAll(r.Context())
		if err != nil {
			log.Println(err.Error())
		}
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(struct{ Status string }{Status: "All users deleted"})
	}
}
