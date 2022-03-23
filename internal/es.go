package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

func createESClient() *elasticsearch.Client {
	esHost := os.Getenv("ES_HOST")
	cfg := elasticsearch.Config{
		Addresses: []string{
			esHost,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	CheckError(err, "Failed to create elasticsearch connection.")
	return es
}

func RemoveOldIndex(indexPrefix string, days int) {
	fmt.Printf("remove index %s older than %d", indexPrefix, days)
	es := createESClient()
	log.Println(es.Info())
}
