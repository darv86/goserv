package feed

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/darv86/goserv/internal/database"
	"github.com/darv86/goserv/internal/utils"
	"github.com/darv86/goserv/shared"
)

func MineDeleteById(apiConf *shared.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("from delete my feed by id router")
		id, err := utils.GetUrlLastParam(r.URL)
		if err != nil {
			log.Println(err.Error())
			return
		}
		params := database.FeedMineDeleteByIdParams{
			ID:     id,
			UserID: apiConf.AuthUser.ID,
		}
		feedDb, err := apiConf.Queries.FeedMineDeleteById(r.Context(), params)
		if err != nil {
			log.Println(err.Error())
		}
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(feedDb)
	}
}
