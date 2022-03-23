package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func getEsHost() string {
	esHost := os.Getenv("ES_HOST")
	if len(esHost) == 0 {
		return "http://localhost:9200"
	}
	return esHost
}

func createESClient() *elasticsearch.Client {
	esHost := getEsHost()
	cfg := elasticsearch.Config{
		Addresses: []string{
			esHost,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	CheckError(err, "Failed to create elasticsearch connection.")
	return es
}

type ESIndex struct {
}

func RemoveOldIndex(indexPrefix string, days int) {
	fmt.Printf("remove index %s older than %d", indexPrefix, days)
	es := createESClient()
	log.Println(es.Info())
	res, err := esapi.CatIndicesRequest{Format: "json"}.Do(context.Background(), es)
	CheckError(err, "Failed to create index request")
	defer res.Body.Close()
	fmt.Println(res.String())
	var r []map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&r)
	CheckError(err, "Failed to parse index response")
	baseDate := time.Now().AddDate(0, 0, -7)
	fmt.Println("base date", baseDate)
	for _, s := range r {
		index := strings.ReplaceAll(fmt.Sprintf("%v", s["index"]), indexPrefix, "")
		fmt.Println(index)
		indexDate, err := time.Parse("2006-01-02", index)
		if err == nil {
			fmt.Println(indexDate.String())
			if indexDate.Before(baseDate) {
				fmt.Println("delete index", index)
			}
		}
	}
}
