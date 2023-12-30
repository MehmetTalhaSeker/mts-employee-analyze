package main

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

type KafkaService struct {
	producer sarama.SyncProducer
}

func NewKafkaService() *KafkaService {
	brokers := []string{"kafka-1:9092"}

	config := sarama.NewConfig()
	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true

	// Producer
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatal(err)
	}

	return &KafkaService{
		producer: producer,
	}
}

func (k *KafkaService) sendWishData(wish *Wish) {
	jsonData, err := json.Marshal(wish)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	msg := &sarama.ProducerMessage{
		Topic: "wish_topic",
		Key:   sarama.StringEncoder(wish.EmployeeID),
		Value: sarama.StringEncoder(jsonData),
	}

	_, _, err = k.producer.SendMessage(msg)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println("Message sent")
}
