package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/zhaoyi0113/es/scheduler/internal"
)

var es *elasticsearch.Client

func beforeAll() {
	log.Println("setup suite")
	esHost := os.Getenv("ES_HOST")
	if len(esHost) == 0 {
		esHost = "http://localhost:9200"
	}
	cfg := elasticsearch.Config{
		Addresses: []string{
			esHost,
		},
	}
	client, err := elasticsearch.NewClient(cfg)
	internal.CheckError(err, "Failed to create es client")
	es = client
	res, err := esapi.IndicesCreateRequest{Index: "aws-log-2020-03-23"}.Do(context.Background(), es)
	internal.CheckError(err, "Failed to create es index")
	fmt.Println(res.String())
}

func TestHelloName(t *testing.T) {
	beforeAll()
	internal.RemoveOldIndex(7)
}
