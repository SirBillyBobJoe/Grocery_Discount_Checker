package api

import (
	"encoding/json"
	"net/http"
)

type SubscriptionPayload struct {
	ItemId string `bson:"_id" json:"itemId"`
	Email  string `json:"email"`
}

type Response struct {
	NameAndAge string
}

type Error struct {
	Code    int
	Message string
}

func writeError(responseWriter http.ResponseWriter, message string, code int) {
	response := Error{
		Code:    code,
		Message: message,
	}

	responseWriter.Header().Set(
		"Content-Type",
		"application/json",
	)
	responseWriter.WriteHeader(code)

	json.NewEncoder(responseWriter).Encode(response)
}

var (
	RequestErrorHandler = func(responseWriter http.ResponseWriter, err error) {
		writeError(responseWriter, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(responseWriter http.ResponseWriter) {
		writeError(responseWriter, "An Unexpected Error occured", http.StatusInternalServerError)
	}
)
