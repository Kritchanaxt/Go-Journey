package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Email   string
	Address Address
}

type Address struct {
	Street  string
	City    string
	Country string
}

func (p Person) Greet() {
	fmt.Printf("Hello, I'm %s, %d years old\n", p.Name, p.Age)
}

func (p *Person) Birthday() {
	p.Age++
}

func main() {
	// Struct initialization
	person1 := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "Bangkok",
			Country: "Thailand",
		},
	}

	person2 := Person{Name: "Jane", Age: 25}

	var person3 Person
	person3.Name = "Bob"
	person3.Age = 35

	// Method calls
	person1.Greet()
	person1.Birthday()
	fmt.Println("After birthday:", person1.Age)

	// Anonymous struct
	student := struct {
		Name  string
		Grade int
	}{
		Name:  "Alice",
		Grade: 95,
	}

	fmt.Println(person1, person2, person3)
	fmt.Println(student)
}