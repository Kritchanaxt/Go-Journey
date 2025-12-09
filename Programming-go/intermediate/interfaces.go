package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}

	printShapeInfo(rect)
	printShapeInfo(circle)

	// Empty interface
	var i interface{}
	i = 42
	fmt.Println(i)
	i = "hello"
	fmt.Println(i)
	i = true
	fmt.Println(i)

	// Type assertion
	var val interface{} = "Hello, Go!"
	str, ok := val.(string)
	if ok {
		fmt.Println("String value:", str)
	}
}