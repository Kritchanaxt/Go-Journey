package main

import "fmt"

func main() {
	// Arithmetic
	a, b := 10, 3
	fmt.Println("Add:", a+b)      // 13
	fmt.Println("Subtract:", a-b) // 7
	fmt.Println("Multiply:", a*b) // 30
	fmt.Println("Divide:", a/b)   // 3
	fmt.Println("Modulus:", a%b)  // 1

	// Comparison
	fmt.Println("Equal:", a == b)         // false
	fmt.Println("Not Equal:", a != b)     // true
	fmt.Println("Greater:", a > b)        // true
	fmt.Println("Less:", a < b)           // false
	fmt.Println("Greater/Equal:", a >= b) // true
	fmt.Println("Less/Equal:", a <= b)    // false

	// Logical
	x, y := true, false
	fmt.Println("AND:", x && y) // false
	fmt.Println("OR:", x || y)  // true
	fmt.Println("NOT:", !x)     // false

	// Bitwise
	fmt.Println("AND:", a&b)          // 2
	fmt.Println("OR:", a|b)           // 11
	fmt.Println("XOR:", a^b)          // 9
	fmt.Println("Left Shift:", a<<1)  // 20
	fmt.Println("Right Shift:", a>>1) // 5

	// Assignment
	c := 10
	c += 5 // c = c + 5
	c -= 3 // c = c - 3
	c *= 2 // c = c * 2
	c /= 4 // c = c / 4
	fmt.Println("Result:", c)
}