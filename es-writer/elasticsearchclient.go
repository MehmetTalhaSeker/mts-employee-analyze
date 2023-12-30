package main

import (
	esv7 "github.com/elastic/go-elasticsearch/v7"
	"log"
)

func newElasticSearchClient() *esv7.Client {
	cfg := esv7.Config{
		Addresses: []string{"http://es01:9200"},
	}

	es, err := esv7.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = res.Body.Close()
	}()

	return es
}
