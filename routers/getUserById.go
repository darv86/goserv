package routers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/darv86/goserv/internal/database"
	"github.com/go-chi/chi/v5"
)

func GetUserByIdRouter(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("from get users by id router")
		//
		param := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			log.Println(err.Error())
			return
		}
		userId := sql.NullInt64{Int64: id, Valid: true}
		users, err := queries.GetUserById(r.Context(), userId)
		if err != nil {
			log.Println(err.Error())
		}
		json.NewEncoder(w).Encode(users)
	}
}
