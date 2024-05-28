package main

import (
	"fmt"
	"net"
	"time"
)

// sendMessages connects to a TCP server at the specified IP and port
// and sends a message every second.
func sendMessages(ip string, port int) {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Message sending loop
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		message := "I am your head if I wasnt already"
		_, err := conn.Write([]byte(message + "\n")) // Ensure to send a newline if the server expects lines
		if err != nil {
			fmt.Println("Error sending message:", err)
			return // Exit if there is an error sending a message
		}
		fmt.Println("Message sent:", message)
	}
}
