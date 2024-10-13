package main

import (
	"log"
	"time"
)

type Notification struct {
	MessageID    int64     `json:"m_id"`
	UserID       string    `json:"c_id"`
	Notification string    `json:"notification"`
	CreatedAt    time.Time `json:"created_at"`
}


func main() {
	// Start the consumer in a separate goroutine
	go func() {
		if err := StartConsumer("alerts"); err != nil {
			log.Fatalf("Consumer error: %v", err)
		}
	}()

	// Start the producer
	if err := StartProducer("alerts"); err != nil {
		log.Fatalf("Producer error: %v", err)
	}
}
