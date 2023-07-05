package router

import (
	"net/http"
	"recipes/api"
	"recipes/internal/handler"

	"github.com/gorilla/mux"
)

func New(h *handler.Handler) http.Handler {
	r := mux.NewRouter()
	r.Use(h.AuthMiddleware)
	r.Use(SetJSONHeader)
	r.HandleFunc("/health", h.HealthCheck).Methods(http.MethodGet)
	// c := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:8000"},
	// 	AllowCredentials: true,
	// })
	// handler := c.Handler(r)
	// api.HandlerFromMux()
	apiRouter := api.HandlerWithOptions(h, api.GorillaServerOptions{BaseRouter: r})
	return apiRouter
}

func SetJSONHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
