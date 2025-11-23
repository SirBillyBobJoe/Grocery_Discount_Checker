package handlers

import (
	"learn_go/internal/app"

	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	App *app.App
}

func RegisterRoutes(router *chi.Mux, app *app.App) {
	router.Use(chimiddle.StripSlashes)

	handler := &Handler{
		App: app,
	}

	router.Route("/api", func(chiRouter chi.Router) {
		chiRouter.Post("/subscribe", handler.subscribe)
		chiRouter.Get("/subscriptions", handler.getSubscriptions)
	})
}
