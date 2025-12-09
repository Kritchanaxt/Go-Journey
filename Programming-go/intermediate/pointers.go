package main

import "fmt"

func modifyValue(val int) {
	val = 100
}

func modifyPointer(ptr *int) {
	*ptr = 100
}

func main() {
	// Pointer basics
	x := 42
	var ptr *int = &x

	fmt.Println("Value:", x)
	fmt.Println("Address:", &x)
	fmt.Println("Pointer:", ptr)
	fmt.Println("Dereference:", *ptr)

	// Modify via pointer
	*ptr = 50
	fmt.Println("Modified x:", x)

	// Pass by value vs pointer
	a := 10
	modifyValue(a)
	fmt.Println("After modifyValue:", a) // 10 (unchanged)

	b := 10
	modifyPointer(&b)
	fmt.Println("After modifyPointer:", b) // 100 (changed)

	// New keyword
	newPtr := new(int)
	*newPtr = 99
	fmt.Println("New pointer value:", *newPtr)
}