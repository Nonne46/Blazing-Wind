package elastic

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

var elasticClient *elasticsearch.Client

// Init .
func Init() {
	dsn := os.Getenv("ELASTIC_URL")

	cfg := elasticsearch.Config{
		Addresses: []string{
			dsn,
		},
	}

	var err error

	elasticClient, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Panicf("Error creating the client: %s", err)
	}

	res, err := elasticClient.Info()
	if err != nil {
		log.Printf("Error getting response: %s", err)
	}

	log.Println(res)
}

// GetDB .
func GetDB() *elasticsearch.Client {
	return elasticClient
}
