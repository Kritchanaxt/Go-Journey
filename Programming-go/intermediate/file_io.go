package main

import (
	"fmt"
	"os"
)

func main() {
	// Write file
	content := []byte("Hello, Go!\nFile I/O Example")
	err := os.WriteFile("example.txt", content, 0644)
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}

	// Read file
	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	fmt.Println("File content:", string(data))

	// Append to file
	file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Open error:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("\nNew line"); err != nil {
		fmt.Println("Write error:", err)
	}
}