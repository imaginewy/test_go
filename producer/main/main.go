package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	broker := "localhost:10004"

	topic := "test"

	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", broker, topic, partition)

	if err != nil {

		fmt.Println("Failed to dial kafka leader:", err)

		return

	}

	defer conn.Close()

	fmt.Println("Connected to Kafka cluster:", broker)
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:10004"},
		Topic:   "test",
	})
	message := kafka.Message{
		Key:   []byte("key"),
		Value: []byte("hello kafka!"),
	}
	err = writer.WriteMessages(context.Background(), message)
	if err != nil {
		fmt.Println("failed to send message,", err)
	}
	fmt.Println("message sent successfully")
}
