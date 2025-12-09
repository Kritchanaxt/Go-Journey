package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func processData(value int) error {
	if value < 0 {
		return &CustomError{Code: 400, Message: "Value must be positive"}
	}
	return nil
}

func main() {
	// Basic error handling
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Result:", result2)
	}

	// Custom error
	err3 := processData(-5)
	if err3 != nil {
		fmt.Println("Custom error:", err3)
	}
}