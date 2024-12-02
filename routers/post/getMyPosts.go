package post

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/darv86/goserv/shared"
)

func GetMyPosts(apiConf *shared.ApiConfig) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		log.Println("from get all my posts router")
		postsDb, err := apiConf.Queries.PostByUser(req.Context(), apiConf.AuthUser.ID)
		if err != nil {
			log.Println(err.Error())
		}
		resp.Header().Add("Content-type", "application/json")
		json.NewEncoder(resp).Encode(postsDb)
	}
}
