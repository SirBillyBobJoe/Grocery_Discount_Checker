package middleware

import (
	"errors"
	"learn_go/api"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var UnauthorizedError = errors.New("Invalid Token")

func Authorization(next http.Handler) http.Handler {

	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		var token string = request.Header.Get("Authorization")

		if token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(responseWriter, UnauthorizedError)
			return
		}

		next.ServeHTTP(responseWriter, request)
	})
}
