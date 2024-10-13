package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var (
	broker = "localhost:9092"
	rng    = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// StartProducer initializes the Kafka producer and produces messages
func StartProducer(topic string) error {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		return fmt.Errorf("failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	// Infinite loop to produce alerts
	for {
		alert := &Notification{
			MessageID:    time.Now().UnixNano(),
			UserID:       fmt.Sprintf("user-%d", rng.Intn(1000)), // Random user ID
			Notification: generateRandomAlert(),
			CreatedAt:    time.Now(),
		}

		// Convert the alert to JSON
		alertJSON, err := json.Marshal(alert)
		if err != nil {
			log.Printf("Error marshaling JSON: %v", err)
			continue
		}

		// Produce the alert to Kafka
		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(fmt.Sprintf("%d", alert.MessageID)),
			Value:          alertJSON,
		}, nil)

		if err != nil {
			log.Printf("Error producing message: %v", err)
		} else {
			LogProducedNotification(alert.Notification, time.Now())
		}

		// Wait for delivery report
		e := <-producer.Events()
		m := e.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			log.Printf("Delivery failed: %v", m.TopicPartition.Error)
		}

		time.Sleep(2 * time.Second) // Pause between each alert
	}
}

// Generate a random alert message
func generateRandomAlert() string {
	alerts := []string{
		"CPU usage is high",
		"Memory usage is at 80%",
		"Disk space is running low",
		"New user signed up",
		"Backup completed successfully",
	}
	return alerts[rng.Intn(len(alerts))]
}
