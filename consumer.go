package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var (
	consumer *kafka.Consumer
)

// StartConsumer initializes and starts the Kafka consumer
func StartConsumer(topic string) error {
	config := &kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          "notification-consumer",
		"auto.offset.reset": "earliest",
	}

	var err error
	consumer, err = kafka.NewConsumer(config)
	if err != nil {
		return err
	}

	// Subscribe to the topic
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return err
	}

	// Consume messages
	defer consumer.Close()
	log.Println("Consumer started, waiting for messages...")

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			var notification Notification
			if err := json.Unmarshal(msg.Value, &notification); err != nil {
				log.Printf("Failed to unmarshal notification: %v", err)
				continue
			}

			LogConsumedNotification(notification.Notification, time.Now())
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
