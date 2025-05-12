package router

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/philippmos/eventsourcinglearninghub/topicservice/internal/handler"
	"github.com/philippmos/eventsourcinglearninghub/topicservice/internal/logger"
)

func ApiRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(RequestIdMiddleware)
	router.Use(LoggingMiddleware)
	router.Use(DefaultHeaderMiddleware)

	router.HandleFunc("/topics", handler.GetTopicsHandler).Methods("GET")
	router.HandleFunc("/topics/{id}", handler.GetTopicByIdHandler).Methods("GET")

	return router
}

func DefaultHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {

		responseWriter.Header().Set("Content-Type", "application/json")

		responseWriter.Header().Set("X-Content-Type-Options", "nosniff")
		responseWriter.Header().Set("X-Frame-Options", "DENY")
		responseWriter.Header().Set("Referrer-Policy", "no-referrer")
		responseWriter.Header().Set("Cache-Control", "no-store")
		responseWriter.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

		next.ServeHTTP(responseWriter, request)
	})
}

func RequestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {

		requestIdHeaderName := "X-Request-Id"

		requestId := request.Header.Get(requestIdHeaderName)

		if requestId == "" {
			requestId = uuid.New().String()
		}

		request.Header.Set(requestIdHeaderName, requestId)
		responseWriter.Header().Set(requestIdHeaderName, requestId)

		next.ServeHTTP(responseWriter, request)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		start := time.Now()

		next.ServeHTTP(responseWriter, request)

		logger.Log.WithFields(map[string]interface{}{
			"method":      request.Method,
			"path":        request.URL.Path,
			"remote_addr": request.RemoteAddr,
			"user_agent":  request.UserAgent(),
			"request_id":  request.Header.Get("X-Request-Id"),
			"duration_ms": time.Since(start).Milliseconds(),
		}).Info("incoming request")
	})
}
