package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting server...")
	InitRedis("localhost:6379")
	startServer()
}
