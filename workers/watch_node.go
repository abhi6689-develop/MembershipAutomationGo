package main

import (
	"context"
	"log"
)

func watchHeadNodes() {
	pubsub := client.Subscribe(context.Background(), "head_nodes_channel")
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(context.Background())
		if err != nil {
			log.Println("Error receiving message:", err)
			continue
		}
		log.Printf("Head node changed: %s\n", msg.Payload)
	}
}
