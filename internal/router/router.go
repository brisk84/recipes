package router

import (
	"net/http"
	"recipes/internal/handler"

	"recipes/api"

	"github.com/gorilla/mux"
)

func New(h *handler.Handler) http.Handler {
	r := mux.NewRouter()
	r.Use(SetJSONHeader)
	r.HandleFunc("/health", h.HealthCheck).Methods(http.MethodGet)
	apiRouter := api.HandlerWithOptions(h, api.GorillaServerOptions{BaseRouter: r})
	return apiRouter
}

func SetJSONHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
