package router

import (
	"net/http"
	"recipes/api"
	"recipes/internal/handler"

	"github.com/gorilla/mux"
)

func New(h *handler.Handler) http.Handler {
	r := mux.NewRouter()
	r.Use(corsMiddleware)
	// r.Use(SetJSONHeader)
	r.Use(h.AuthMiddleware)
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

// func SetJSONHeader(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Add("Content-Type", "application/json")
// 		next.ServeHTTP(w, r)
// 	})
// }

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		// w.Header().Add("Access-Control-Allow-Credentials", "true")
		// w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("content-type", "application/json;charset=UTF-8")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
