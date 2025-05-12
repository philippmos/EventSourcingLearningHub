package main

import (
	"log"
	"net/http"

	"github.com/philippmos/eventsourcinglearninghub/topicservice/internal/router"
)

func main() {
	apiRouter := router.ApiRouter()

	log.Println("Starting server on :5002...")
	log.Fatal(http.ListenAndServe(":5002", apiRouter))
}
