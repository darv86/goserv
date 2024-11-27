package feed

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/darv86/goserv/internal/database"
	"github.com/darv86/goserv/shared"
)

func Create(apiConfig *shared.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("from feed create router")
		var params database.FeedCreateParams
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			log.Println(err.Error())
			return
		}
		params.UserID = apiConfig.AuthUser.ID
		feedDb, err := apiConfig.Queries.FeedCreate(r.Context(), params)
		if err != nil {
			log.Println(err.Error())
			return
		}
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(feedDb)
	}
}
