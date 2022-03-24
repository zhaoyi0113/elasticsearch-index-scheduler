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

func RemoveOldIndex(prefix string, days int) {
	fmt.Printf("remove index older than %d", days)
	es := createESClient()
	log.Println(es.Info())
	res, err := esapi.CatIndicesRequest{Format: "json", H: []string{"i", "creation.date.string"}}.Do(context.Background(), es)
	CheckError(err, "Failed to create index request")
	defer res.Body.Close()
	fmt.Println(res.String())
	var r []map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&r)
	CheckError(err, "Failed to parse index response")
	baseDate := time.Now().AddDate(0, 0, -days)
	for _, s := range r {
		index := fmt.Sprintf("%v", s["i"])
		indexCreatedDate := fmt.Sprintf("%v", s["creation.date.string"])
		indexDate, err := time.Parse(time.RFC3339, indexCreatedDate)
		if err == nil && strings.HasPrefix(index, prefix) {
			fmt.Println(index, indexDate.String())
			if indexDate.Before(baseDate) {
				fmt.Println("delete index", index)
				res, err := esapi.IndicesDeleteRequest{Index: []string{index}}.Do(context.Background(), es)
				CheckError(err, "Failed to delete index:"+index)
				fmt.Println("delete index success" + index + "," + res.String())
			}
		} else {
			log.Println("Failed to parse index", err)
		}
	}
}
