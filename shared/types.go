package shared

import (
	"net/http"
	"time"

	"github.com/darv86/goserv/internal/database"
)

type IRouter interface {
	Get(pattern string, handlerFn http.HandlerFunc)
	Post(pattern string, handlerFn http.HandlerFunc)
	Delete(pattern string, handlerFn http.HandlerFunc)
}

type IAuthUser interface {
	database.User | []database.User
}

type ApiConfig struct {
	Queries  *database.Queries
	AuthUser database.User
	Router   IRouter
}

type ScraperConfig struct {
	Queries        *database.Queries
	TickInterval   time.Duration
	MaxFeedsAtTime int
}
