package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"learn_go/api"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func (handler *Handler) subscribe(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	ctx := httpRequest.Context()

	var request api.SubscriptionPayload

	if err := json.NewDecoder(httpRequest.Body).Decode(&request); err != nil {
		api.RequestErrorHandler(responseWriter, errors.New("bad request"))
		return
	}

	shortCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := handler.App.SubscriptionRepository.SaveSubscription(shortCtx, request); err != nil {
		log.Error(err)
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusCreated)
}

func (handler *Handler) getSubscriptions(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	ctx := httpRequest.Context()

	shortCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	handler.App.SubscriptionRepository.GetAllSubscriptions(shortCtx)
	if subscriptions, err := handler.App.SubscriptionRepository.GetAllSubscriptions(shortCtx); err != nil {
		log.Error(err)

		responseWriter.WriteHeader(http.StatusNotFound)
		return
	} else {
		json.NewEncoder(responseWriter).Encode(subscriptions)

		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusOK)
	}
}
