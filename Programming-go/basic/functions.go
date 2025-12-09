package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}

func calculate(a, b int) (sum int, diff int, prod int) {
	sum = a + b
	diff = a - b
	prod = a * b
	return // Named return
}

func variadicSum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	result := add(5, 3)
	fmt.Println("Add:", result)

	x, y := swap("hello", "world")
	fmt.Println("Swap:", x, y)

	sum, diff, prod := calculate(10, 5)
	fmt.Println("Sum:", sum, "Diff:", diff, "Prod:", prod)

	total := variadicSum(1, 2, 3, 4, 5)
	fmt.Println("Variadic Sum:", total)

	// Anonymous function
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Println("Anonymous:", multiply(4, 5))

	// Closure
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
}