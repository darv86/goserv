package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/darv86/goserv/internal/database"
	"github.com/darv86/goserv/internal/utils"
	"github.com/go-chi/chi/v5"
)

func GetById(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("from get users by id router")
		param := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			log.Println(err.Error())
			return
		}
		// userId := sql.NullInt64{Int64: id, Valid: true}
		userDb, err := queries.UserGetById(r.Context(), id)
		if err != nil {
			log.Println(err.Error())
		}
		user := utils.GetStructTypeOf[User](userDb)
		json.NewEncoder(w).Encode(user)
	}
}
