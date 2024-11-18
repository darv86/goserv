package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/darv86/goserv/internal/database"
)

func GetAll(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("from get users router")
		usersDb, err := queries.GetAll(r.Context())
		if err != nil {
			log.Println(err.Error())
		}
		//
		var users []User
		for _, userDb := range usersDb {
			user := User{
				ID:        int(userDb.ID.Int64),
				CreatedAt: userDb.CreatedAt,
				UpdatedAt: userDb.UpdatedAt,
				Name:      userDb.Name,
			}
			users = append(users, user)
		}
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}
