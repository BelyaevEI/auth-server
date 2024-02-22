package storage

import (
	"context"

	"github.com/BelyaevEI/auth-server/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storager interface {
}

type mongoDB struct {
	collection *mongo.Collection
}

func NewConnect(cfg config.Config) (Storager, error) {

	ctx := context.Background()

	// Connect to mongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.DSN))
	if err != nil {
		return mongoDB{}, nil
	}

	collction := client.Database("user_token").Collection("tokens")

	return mongoDB{collection: collction}, nil
}
