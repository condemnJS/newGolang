package transport

import (
	"github.com/gorilla/mux"
	"net/http"
	"example-api-project/pkg/controllers"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/health", controllers.HealthHandler).Methods(http.MethodGet)
	return r
}
