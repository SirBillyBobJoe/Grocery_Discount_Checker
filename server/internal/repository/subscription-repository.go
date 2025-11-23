package repository

import (
	"context"
	"learn_go/api"
	"learn_go/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SubcriptionRepository struct {
	Collection *mongo.Collection
}

func NewSubscriptionRepository(database *mongo.Database) *SubcriptionRepository {
	return &SubcriptionRepository{
		Collection: database.Collection("subscriptions"),
	}
}

func (repository *SubcriptionRepository) SaveSubscription(ctx context.Context, subscription api.SubscriptionPayload) error {
	filter := bson.M{"_id": subscription.ItemId}
	update := bson.M{"$addToSet": bson.M{"emails": subscription.Email}, "$setOnInsert": bson.M{"orignalPrice": 0.0, "currentPrice": 0.0}}
	opts := options.Update().SetUpsert(true)

	_, err := repository.Collection.UpdateOne(ctx, filter, update, opts)

	return err
}

func (repository *SubcriptionRepository) UpdateSubscription(ctx context.Context, subscription model.SubcriptionModel, originalPrice float32, currentPrice float32) error {
	filter := bson.M{"_id": subscription.ItemId}
	update := bson.M{"$set": bson.M{"orignalPrice": originalPrice, "currentPrice": currentPrice}}
	opts := options.Update().SetUpsert(true)

	_, err := repository.Collection.UpdateOne(ctx, filter, update, opts)

	return err
}

func (repository *SubcriptionRepository) GetAllSubscriptions(ctx context.Context) ([]model.SubcriptionModel, error) {
	filter := bson.M{}
	cursor, err := repository.Collection.Find(ctx, filter)

	var subscriptions []model.SubcriptionModel

	cursor.All(ctx, &subscriptions)

	return subscriptions, err
}
