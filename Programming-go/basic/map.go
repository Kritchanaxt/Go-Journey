package main

import "fmt"

func main() {
	// Map Declaration
	var map1 map[string]int
	map1 = make(map[string]int)

	map2 := make(map[string]string)
	map3 := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 7,
	}

	// Map Operations
	map2["name"] = "John"
	map2["city"] = "Bangkok"

	// Check if key exists
	value, exists := map3["apple"]
	if exists {
		fmt.Println("Apple count:", value)
	}

	// Delete key
	delete(map3, "banana")

	// Iterate map
	for key, value := range map3 {
		fmt.Printf("%s: %d\n", key, value)
	}

	fmt.Println(map1, map2, map3)
}