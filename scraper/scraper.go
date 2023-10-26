package scraper

import (
	tlsclient "github.com/bogdanfinn/tls-client"
	"github.com/jakubje/property-aggregator/util"
	"go.mongodb.org/mongo-driver/mongo"
)

//type PropertyScraper interface {
//	Start() error
//}

type PropertyScraperProcessor struct {
	config          util.Config
	client          tlsclient.HttpClient
	mongoCollection *mongo.Collection
}

func NewPropertyScraper(config util.Config, client *tlsclient.HttpClient, mongoCollection *mongo.Collection) *PropertyScraperProcessor {
	return &PropertyScraperProcessor{
		config:          config,
		client:          *client,
		mongoCollection: mongoCollection,
	}
}
func (s *PropertyScraperProcessor) Start() error {
	//go s.getRightMoveData()
	go s.zooplaInit()

	return nil
}
