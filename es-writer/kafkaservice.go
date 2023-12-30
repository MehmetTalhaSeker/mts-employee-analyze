package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	esv7 "github.com/elastic/go-elasticsearch/v7"
	"log"
)

type kafkaService struct {
	consumer *kafka.Consumer
	esc      *esv7.Client
}

func newKafkaService(esc *esv7.Client) *kafkaService {
	config := kafka.ConfigMap{
		"bootstrap.servers":  "kafka-1:9092",
		"group.id":           "groupID",
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
	}

	cs, err := kafka.NewConsumer(&config)
	if err != nil {
		log.Fatal(err)
	}

	if err := cs.Subscribe("wish_analysis_topic", nil); err != nil {
		log.Fatal(err)
	}

	return &kafkaService{
		consumer: cs,
		esc:      esc,
	}
}

// TODO: actor model implementation
func (s *kafkaService) startProcessing() {
	for {
		ev := s.consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			// flaw-1
			go func() {
				// flaw-2
				_, err := s.consumer.CommitMessage(e)
				if err != nil {
					fmt.Printf("ERROR: %v", err)
					return
				}

				sit := elasticsearchMsg{
					Word:  string(e.Key),
					Count: string(e.Value),
				}

				data, err := json.Marshal(sit)
				if err != nil {
					log.Printf("Error marshaling document: %s", err)
					return
				}

				// flaw-3
				_, err = s.esc.Index("systemindex", bytes.NewReader(data))
				if err != nil {
					log.Printf("Error ES Indexing document: %s", err)
				}
			}()
		}
	}
}

type elasticsearchMsg struct {
	Word  string `json:"word"`
	Count string `json:"count"`
}
