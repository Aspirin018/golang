package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	message := make(chan int, 10)
	//consumer
	go func() {
		//time.Sleep(1 * time.Second)
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			fmt.Printf("first received message: %d\n\n", <-message)
		}
	}()

	go func() {
		//time.Sleep(1 * time.Second)
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			fmt.Printf("second received message: %d\n\n", <-message)
		}
	}()

	//producer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			random := rand.Intn(100)
			message <- random
			fmt.Printf("send message: %d\n", random)
		}
	}()

	time.Sleep(20 * time.Second)
	fmt.Println("main exit!")
}
