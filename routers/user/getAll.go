package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/darv86/goserv/internal/database"
	"github.com/darv86/goserv/internal/utils"
)

func GetAll(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("from get users router")
		usersDb, err := queries.UserGetAll(r.Context())
		if err != nil {
			log.Println(err.Error())
		}
		var users []User
		for _, userDb := range usersDb {
			user := utils.GetStructTypeOf[User](userDb)
			users = append(users, user)
		}
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}
