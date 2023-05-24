package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	messages := make(chan int, 10)
	defer close(messages)
	// consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			fmt.Printf("consumer read message: %d\n", <-messages)
		}
	}()

	// producer
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		fmt.Println("producer send message...")
		messages <- rand.Int()
	}
}
