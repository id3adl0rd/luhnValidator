package route

import (
	"cardValidator/pkg/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRoute() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/validate", cardHandler).Methods(http.MethodGet)
	r.Use(middleware.LoggingMiddleware, middleware.ContentTypeMiddleware, middleware.RecoveryMiddleware)

	return r
}
