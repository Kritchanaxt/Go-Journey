package main

import "fmt"

func main() {
	// If-Else
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// If with short statement
	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even")
	}

	// Switch
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other day")
	}

	// Switch without condition
	hour := 14
	switch {
	case hour < 12:
		fmt.Println("Good Morning")
	case hour < 18:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}

	// Type Switch
	var i interface{} = "hello"
	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case bool:
		fmt.Println("Boolean:", v)
	default:
		fmt.Println("Unknown type")
	}
}