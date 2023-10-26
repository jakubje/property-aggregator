package db

import (
	"context"
	"github.com/jakubje/property-aggregator/util"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewConnection(config util.Config) (*mongo.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(config.DBMongoUri)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to mongo")
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Info().Msg("connected to mongo")

	return client, nil
}

func NewCollection(config util.Config, client *mongo.Client) (*mongo.Collection, error) {

	db := client.Database(config.DBName)
	collection := db.Collection(config.CollectionName)

	return collection, nil
}
