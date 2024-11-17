package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/darv86/goserv/internal/database"
)

func GetUsersRouter(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("from get users router")
		users, err := queries.GetUsers(r.Context())
		if err != nil {
			log.Println(err.Error())
		}
		json.NewEncoder(w).Encode(users)
	}
}
