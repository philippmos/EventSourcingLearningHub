package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/philippmos/eventsourcinglearninghub/topicservice/internal/handler"
)

func ApiRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(DefaultHeaderMiddleware)

	router.HandleFunc("/topics", handler.GetTopicsHandler).Methods("GET")
	router.HandleFunc("/topics/{id}", handler.GetTopicByIdHandler).Methods("GET")

	return router
}

func DefaultHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

		next.ServeHTTP(w, r)
	})
}
