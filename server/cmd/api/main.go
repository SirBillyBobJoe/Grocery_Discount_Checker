package main

import (
	"fmt"
	"learn_go/internal/app"
	handlers "learn_go/internal/handlers"
	jobs "learn_go/internal/jobs"
	"learn_go/internal/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	log "github.com/sirupsen/logrus"

	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	godotenv.Load()

	rootCtx := context.Background()
	log.SetReportCaller(true)

	ctx, cancel := context.WithTimeout(rootCtx, 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	defer func() {
		disconnectCtx, cancel := context.WithTimeout(rootCtx, 5*time.Second)
		defer cancel()
		if err := client.Disconnect(disconnectCtx); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Connection to MongoDB closed.")
		}
	}()

	var router *chi.Mux = chi.NewRouter()

	database := client.Database("discounts")
	subscriptionRepository := repository.NewSubscriptionRepository(database)
	App := app.NewApp(subscriptionRepository)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	handlers.RegisterRoutes(router, App)

	fmt.Println("Starting Go API Server")

	jobCtx, jobCancel := context.WithCancel(rootCtx)
	defer jobCancel()
	go jobs.SubscriptionJob(jobCtx, App)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Error("Server error: ", err)
	}
}
