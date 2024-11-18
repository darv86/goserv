package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/darv86/goserv/internal/database"
)

func Create(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("from user create router")
		//
		var params database.CreateParams
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			log.Println(err.Error())
			return
		}
		//
		userDb, err := queries.Create(r.Context(), params)
		if err != nil {
			log.Println(err.Error())
			return
		}
		user := User{
			ID:        int(userDb.ID.Int64),
			CreatedAt: userDb.CreatedAt,
			UpdatedAt: userDb.UpdatedAt,
			Name:      userDb.Name,
		}
		//
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}
