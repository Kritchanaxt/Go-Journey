package main

import "fmt"

func main() {
	// Array (ขนาดคงที่)
	var arr1 [3]int = [3]int{1, 2, 3}
	arr2 := [5]string{"a", "b", "c", "d", "e"}
	arr3 := [...]int{10, 20, 30, 40}

	// Slice (ขนาดยืดหยุ่น)
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]string, 3, 5) // length 3, capacity 5
	slice3 := arr2[1:4]            // slice from array

	// Slice Operations
	slice1 = append(slice1, 6, 7, 8)
	slice4 := []int{100, 200, 300}
	slice1 = append(slice1, slice4...)
	length := len(slice1)
	capacity := cap(slice1)

	// 2D Array/Slice
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(slice1, slice2, slice3)
	fmt.Println("Length:", length, "Capacity:", capacity)
	fmt.Println(matrix)
}