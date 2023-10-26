package main

import (
	"context"
	tlsclient "github.com/bogdanfinn/tls-client"
	"github.com/jakubje/property-aggregator/api"
	"github.com/jakubje/property-aggregator/db"
	"github.com/jakubje/property-aggregator/scraper"
	"github.com/jakubje/property-aggregator/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config:")
	}
	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	mongoClient, err := db.NewConnection(config)
	if err != nil {
		log.Fatal().Err(err).Msg("error creating mongo client client: ")
	}
	defer mongoClient.Disconnect(context.Background())
	mongoCollection, err := db.NewCollection(config, mongoClient)
	if err != nil {
		log.Fatal().Err(err).Msg("error getting collection: ")
	}

	client, err := util.NewClient()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot get client: ")
	}
	go runScraper(config, client, mongoCollection)
	runGinServer(config, mongoCollection)

	//time.Sleep(15 * time.Second)
}

func runScraper(config util.Config, client *tlsclient.HttpClient, mongoCollection *mongo.Collection) {

	propertyScraper := scraper.NewPropertyScraper(config, client, mongoCollection)
	log.Info().Msg("start property scraper")
	err := propertyScraper.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start crypto processor")
	}
}

func runGinServer(config util.Config, mongoCollection *mongo.Collection) {
	server, err := api.NewServer(config, mongoCollection)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server: ")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server: ")
	}
}
