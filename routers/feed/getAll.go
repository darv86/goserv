package feed

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/darv86/goserv/shared"
)

func GetAll(apiConf *shared.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("from get feeds router")
		feedsDb, err := apiConf.Queries.FeedGetAll(r.Context())
		if err != nil {
			log.Println(err.Error())
		}
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(feedsDb)
	}
}
