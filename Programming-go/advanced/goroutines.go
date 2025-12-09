package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("Letter: %c\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func sum(a, b int, c chan int) {
	c <- a + b
}

func main() {
	// Basic goroutine
	go printNumbers()
	go printLetters()
	time.Sleep(600 * time.Millisecond)

	// Channels
	ch := make(chan int)
	go sum(10, 20, ch)
	result := <-ch
	fmt.Println("Sum from channel:", result)

	// Buffered channel
	messages := make(chan string, 2)
	messages <- "Hello"
	messages <- "World"
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// Channel direction
	onlySend := make(chan<- int)
	onlyReceive := make(<-chan int)
	_ = onlySend
	_ = onlyReceive

	// Select statement
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Channel 1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}
}