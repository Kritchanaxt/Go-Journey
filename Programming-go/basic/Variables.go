package main

import "fmt"

func main() {
	// Manual Declaration
	var name string = "Bak Wave"
	var age int = 25
	var salary float64 = 50000.50
	var isActive bool = true

	// Type Inference
	city := "Nong Bua Lamphu"
	score := 95.5

	// Multiple Variables
	var a, b, c int = 1, 2, 3
	x, y, z := "Go", 3.14, true

	// Constants
	const PI = 3.14159
	const AppName = "MyApp"

	// Zero Values
	var defaultInt int       // 0
	var defaultFloat float64 // 0.0
	var defaultString string // ""
	var defaultBool bool     // false

	fmt.Println(name, age, salary, isActive, city, score)
	fmt.Println(a, b, c, x, y, z)
	fmt.Println(PI, AppName)
	fmt.Println(defaultInt, defaultFloat, defaultString, defaultBool)
}