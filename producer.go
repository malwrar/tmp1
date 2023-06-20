package main

import (
	"context"
	"log"
  "time"
  "fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println("Sending message...")

	    w := &kafka.Writer{
	    	Addr:     kafka.TCP("demo-kafka-0.demo-kafka-headless.demo.svc.cluster.local:9092"),
	    	Topic:    "demo-topic",
	    	Balancer: &kafka.LeastBytes{},
	    }

	    err := w.WriteMessages(context.Background(),
	    	kafka.Message{Value: []byte("Hello Kafka!")},
	    )
	    if err != nil {
	    	log.Fatal("Failed to send message:", err)
	    }

	    if err := w.Close(); err != nil {
	    	log.Fatal("Failed to close writer:", err)
	    }
		}
	}
}
