package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	messageReceived := make(chan bool)

	go func() {
		for scanner.Scan() {
			text := scanner.Text()
			log.Printf("Received: %s", text)
			messageReceived <- true
			addr := conn.RemoteAddr().(*net.TCPAddr).IP.String()
			RefreshRedisKey(addr, 20*time.Second)
		}
		if err := scanner.Err(); err != nil {
			log.Println("Error reading:", err)
		}
	}()

	for {
		select {
		case <-messageReceived:
			// Reset the orphan variable if a message is received
			SetOrphan(false)
		case <-time.After(10 * time.Second):
			SetOrphan(true)
			return
		}
	}
}

func startServer() {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Println("Server started on :9999")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
