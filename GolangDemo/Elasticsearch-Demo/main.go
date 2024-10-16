package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200/"},
		// CloudID: "My_deployment:YXNpYS1zb3V0aGVhc3QxLmdjcC5lbGFzdGljLWNsb3VkLmNvbSRhMTI3N2M2YmI4YjI0MGYxYmUwNDk4MTlkNWQ2YTczMyQ3NzRhY2I0N2Q5OTg0YjBmOWZjYjE4NThiODc4MDdiMQ==",
		// APIKey:  "N0JZSmg1SUJIZlpiNDFVZEtXM3M6MllxYVdDRFpSakNfZlVUX01zNldsQQ==",
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// API Key should have cluster monitoring rights
	_, err = es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	indexName := "cpe-000001"
	// count(indexName, es)
	update(indexName, es)
}

func count(indexName string, es *elasticsearch.Client) {
	req := esapi.CountRequest{
		Index: []string{indexName},
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting count: %s", err)
	}
	defer res.Body.Close()
	// Check for a successful response
	if res.IsError() {
		log.Fatalf("[%s] Error: %s", res.Status(), res.String())
	}

	// Parse the response
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	fmt.Println(r)
}

func update(indexName string, es *elasticsearch.Client) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"script": map[string]interface{}{
			"source": "ctx._source['Deprecated'] = 'true'", // Script to update the field 'status' to 'active'
		},
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"CpeNameId": "EF276652-72B8-4057-B229-512DE9AEA695", // Query to match documents where 'status' equals 'inactive'
			},
		},
	}
	// Encode the query to JSON
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	req := esapi.UpdateByQueryRequest{
        Index: []string{indexName},          // Replace with your index name
        Body:  &buf,                          // The query body
    }

	 // Send request
	 res, err := req.Do(context.Background(), es)
	 if err != nil {
		 log.Printf("Error getting response: %s", err)
	 }
	 defer res.Body.Close()
	 fmt.Println(res)
	 
}

