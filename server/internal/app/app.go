package app

import (
	"learn_go/internal/repository"
)

type App struct {
	SubscriptionRepository *repository.SubcriptionRepository
}

func NewApp(subscriptionRepository *repository.SubcriptionRepository) *App {
	return &App{
		SubscriptionRepository: subscriptionRepository,
	}
}
