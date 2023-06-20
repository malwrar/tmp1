package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"demo-kafka-0.demo-kafka-headless.demo.svc.cluster.local:9092"},
		GroupID:   "demo-group",
		Topic:     "demo-topic",
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Failed to read message:", err)
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("Failed to close reader:", err)
	}
}
