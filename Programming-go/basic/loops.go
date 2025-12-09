package main

import "fmt"

func main() {
	// For loop (traditional)
	for i := 0; i < 5; i++ {
		fmt.Println("Count:", i)
	}

	// For loop (while style)
	j := 0
	for j < 5 {
		fmt.Println("While:", j)
		j++
	}

	// Infinite loop
	/*
		count := 0
		for {
			fmt.Println("Infinite loop:", count)
			count++
			if count >= 3 {
				break
			}
		}
	*/

	// For range (array/slice)
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// For range (map)
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	for key, value := range colors {
		fmt.Printf("%s: %s\n", key, value)
	}

	// For range (string)
	for index, char := range "Hello" {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}

	// Break and Continue
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue // Skip 3
		}
		if i == 7 {
			break // Stop at 7
		}
		fmt.Println("Number:", i)
	}
}