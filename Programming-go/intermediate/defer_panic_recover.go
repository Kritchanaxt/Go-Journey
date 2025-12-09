package main

import "fmt"

func deferExample() {
	defer fmt.Println("This is executed last")
	defer fmt.Println("This is executed second")
	fmt.Println("This is executed first")

	// Defer with loop
	for i := 0; i < 3; i++ {
		defer fmt.Println("Loop defer:", i)
	}
}

func mayPanic() {
	panic("Something went wrong!")
}

func recoverExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("Before panic")
	mayPanic()
	fmt.Println("After panic") // This won't execute
}

func main() {
	fmt.Println("Defer Example:")
	deferExample()

	fmt.Println("\nRecover Example:")
	recoverExample()
}