package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/philippmos/eventsourcinglearninghub/topicservice/internal/model"
)

func GetTopicsHandler(responseWriter http.ResponseWriter, request *http.Request) {
	topics := model.GetAllTopics()
	json.NewEncoder(responseWriter).Encode(topics)
}

func GetTopicByIdHandler(responseWriter http.ResponseWriter, request *http.Request) {
	requestValues := mux.Vars(request)

	id, idParseError := strconv.Atoi(requestValues["id"])

	if idParseError != nil {
		http.Error(responseWriter, "Invalid topic id", http.StatusBadRequest)
		return
	}

	topic, found := model.GetTopicById(id)

	if !found {
		http.Error(responseWriter, "Topic not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(responseWriter).Encode(topic)
}
