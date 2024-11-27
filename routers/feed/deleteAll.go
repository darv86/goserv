package feed

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/darv86/goserv/shared"
)

func DeleteAll(apiConf *shared.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("from delete feeds router")
		err := apiConf.Queries.FeedDeleteAll(r.Context())
		if err != nil {
			log.Println(err.Error())
		}
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(struct{ Status string }{Status: "All feeds deleted"})
	}
}
