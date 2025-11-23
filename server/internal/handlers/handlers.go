package handlers

import (
	"learn_go/internal/app"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	App *app.App
}

func RegisterRoutes(router *chi.Mux, app *app.App) {
	handler := &Handler{
		App: app,
	}

	router.Route("/api", func(chiRouter chi.Router) {
		chiRouter.Post("/subscribe", handler.subscribe)
		chiRouter.Get("/subscriptions", handler.getSubscriptions)
	})
}
