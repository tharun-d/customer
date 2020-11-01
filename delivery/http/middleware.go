package http

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/rs/cors"
)

const (
	cacheMaxAge = 86400
)

func CORS(handler http.Handler) http.Handler {
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           cacheMaxAge,
	})

	return corsHandler.Handler(handler)
}

func Recover(handler http.Handler) http.Handler {
	return handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)
}

func DefaultHandler(handler http.Handler) http.Handler {
	return CORS(handlers.CompressHandler(Recover(handler)))
}
