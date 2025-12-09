package main

import (
	"fmt"
	"time"
)

// ==================== 1. ตัวแปร (Variables) ====================

func variablesExample() {
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

// ==================== 2. ชนิดข้อมูล (Data Types) ====================

func dataTypesExample() {
	// Numbers
	var i8 int8 = 127     // -128 to 127
	var i16 int16 = 32767 // -32768 to 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var u8 uint8 = 255 // 0 to 255
	var f32 float32 = 3.14
	var f64 float64 = 3.14159265359

	// Strings
	var str string = "Hello, Go!"
	multiline := `This is
	a multiline
	string`

	// Boolean
	var isTrue bool = true
	var isFalse bool = false

	// Byte & Rune
	var byteVal byte = 'A' // uint8
	var runeVal rune = '☺' // int32, Unicode

	fmt.Println(i8, i16, i32, i64, u8, f32, f64)
	fmt.Println(str, multiline)
	fmt.Println(isTrue, isFalse)
	fmt.Println(byteVal, runeVal)
}

// ==================== 3. Array & Slice ====================

func arraySliceExample() {
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

// ==================== 4. Map ====================

func mapExample() {
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

// ==================== 5. Operators ====================

func operatorsExample() {
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

// ==================== 6. Control Flow ====================

func controlFlowExample() {
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

// ==================== 7. Loops ====================

func loopsExample() {
	// For loop (traditional)
	for i := 0; i < 5; i++ {
		fmt.Println("Count:", i)
	}

	// For loop (while style)
	j := 0
	for j < 5 {
		fmt.Println("While:", j)
		j++
	}

	// Infinite loop
	/*
		count := 0
		for {
			fmt.Println("Infinite loop:", count)
			count++
			if count >= 3 {
				break
			}
		}
	*/

	// For range (array/slice)
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// For range (map)
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	for key, value := range colors {
		fmt.Printf("%s: %s\n", key, value)
	}

	// For range (string)
	for index, char := range "Hello" {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}

	// Break and Continue
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue // Skip 3
		}
		if i == 7 {
			break // Stop at 7
		}
		fmt.Println("Number:", i)
	}
}

// ==================== 8. Functions ====================

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

func functionsExample() {
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

// ==================== 9. Pointers ====================

func modifyValue(val int) {
	val = 100
}

func modifyPointer(ptr *int) {
	*ptr = 100
}

func pointersExample() {
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

// ==================== 10. Structs ====================

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

func structsExample() {
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

// ==================== 11. Interfaces ====================

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

func interfacesExample() {
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

// ==================== 12. Error Handling ====================

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

func errorHandlingExample() {
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

// ==================== 13. Goroutines & Channels ====================

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("Letter: %c\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func sum(a, b int, c chan int) {
	c <- a + b
}

func goroutinesExample() {
	// Basic goroutine
	go printNumbers()
	go printLetters()
	time.Sleep(600 * time.Millisecond)

	// Channels
	ch := make(chan int)
	go sum(10, 20, ch)
	result := <-ch
	fmt.Println("Sum from channel:", result)

	// Buffered channel
	messages := make(chan string, 2)
	messages <- "Hello"
	messages <- "World"
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// Channel direction
	onlySend := make(chan<- int)
	onlyReceive := make(<-chan int)
	_ = onlySend
	_ = onlyReceive

	// Select statement
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Channel 1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}
}

// ==================== 14. Defer, Panic, Recover ====================

func deferExample() {
	defer fmt.Println("This is executed last")
	defer fmt.Println("This is executed second")
	fmt.Println("This is executed first")

	// Defer with loop
	for i := 0; i < 3; i++ {
		defer fmt.Println("Loop defer:", i)
	}
}

func mayPanic() {
	panic("Something went wrong!")
}

func recoverExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("Before panic")
	mayPanic()
	fmt.Println("After panic") // This won't execute
}

// ==================== 15. File I/O ====================

func fileIOExample() {
	// Write file
	content := []byte("Hello, Go!\nFile I/O Example")
	err := fmt.Errorf("file operations skipped in example")
	_ = content
	_ = err
	// err := os.WriteFile("example.txt", content, 0644)
	// if err != nil {
	// 	fmt.Println("Write error:", err)
	// 	return
	// }

	// Read file
	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	fmt.Println("Read error:", err)
	// 	return
	// }
	// fmt.Println("File content:", string(data))

	// Append to file
	// file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	fmt.Println("Open error:", err)
	// 	return
	// }
	// defer file.Close()

	// if _, err := file.WriteString("\nNew line"); err != nil {
	// 	fmt.Println("Write error:", err)
	// }
}

// ==================== 16. JSON ====================

type User struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Active   bool     `json:"is_active"`
	Tags     []string `json:"tags,omitempty"`
}

func jsonExample() {
	// Marshal (struct to JSON)
	user := User{
		ID:       1,
		Username: "john_doe",
		Email:    "john@example.com",
		Active:   true,
		Tags:     []string{"admin", "developer"},
	}

	fmt.Printf("User created: %+v\n", user)

	_ = fmt.Errorf("json operations skipped in example")

	// jsonData, err := json.Marshal(user)
	// if err != nil {
	// 	fmt.Println("Marshal error:", err)
	// 	return
	// }
	// fmt.Println("JSON:", string(jsonData))

	// Marshal with indent
	// prettyJSON, _ := json.MarshalIndent(user, "", "  ")
	// fmt.Println("Pretty JSON:\n", string(prettyJSON))

	// Unmarshal (JSON to struct)
	// jsonStr := `{"id":2,"username":"jane_doe","email":"jane@example.com","is_active":true}`
	// var user2 User
	// err = json.Unmarshal([]byte(jsonStr), &user2)
	// if err != nil {
	// 	fmt.Println("Unmarshal error:", err)
	// 	return
	// }
	// fmt.Printf("Unmarshaled: %+v\n", user2)
}

// ==================== 17. GIN FRAMEWORK ====================

// Product struct สำหรับ Gin
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []Product{
	{ID: 1, Name: "Laptop", Price: 999.99},
	{ID: 2, Name: "Mouse", Price: 29.99},
	{ID: 3, Name: "Keyboard", Price: 79.99},
}

func ginExample() {
	// สร้าง Gin router
	router := gin.Default()

	// Middleware
	router.Use(func(c *gin.Context) {
		fmt.Println("Request:", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	// GET - ดึงข้อมูลทั้งหมด
	router.GET("/products", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":   "success",
			"products": products,
		})
	})

	// GET - ดึงข้อมูลตาม ID
	router.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, product := range products {
			if fmt.Sprintf("%d", product.ID) == id {
				c.JSON(200, product)
				return
			}
		}
		c.JSON(404, gin.H{"error": "Product not found"})
	})

	// POST - เพิ่มข้อมูล
	router.POST("/products", func(c *gin.Context) {
		var newProduct Product
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		newProduct.ID = len(products) + 1
		products = append(products, newProduct)
		c.JSON(201, newProduct)
	})

	// PUT - แก้ไขข้อมูล
	router.PUT("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedProduct Product
		if err := c.ShouldBindJSON(&updatedProduct); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		for i, product := range products {
			if fmt.Sprintf("%d", product.ID) == id {
				products[i].Name = updatedProduct.Name
				products[i].Price = updatedProduct.Price
				c.JSON(200, products[i])
				return
			}
		}
		c.JSON(404, gin.H{"error": "Product not found"})
	})

	// DELETE - ลบข้อมูล
	router.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, product := range products {
			if fmt.Sprintf("%d", product.ID) == id {
				products = append(products[:i], products[i+1:]...)
				c.JSON(200, gin.H{"message": "Product deleted"})
				return
			}
		}
		c.JSON(404, gin.H{"error": "Product not found"})
	})

	// Query Parameters
	router.GET("/search", func(c *gin.Context) {
		name := c.Query("name")
		minPrice := c.DefaultQuery("min_price", "0")
		c.JSON(200, gin.H{
			"search": name,
			"min":    minPrice,
		})
	})

	// Group Routes
	api := router.Group("/api/v1")
	{
		api.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{"users": []string{"user1", "user2"}})
		})
		api.GET("/orders", func(c *gin.Context) {
			c.JSON(200, gin.H{"orders": []string{"order1", "order2"}})
		})
	}

	// Static Files
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./favicon.ico")

	// HTML Template
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Home Page",
		})
	})

	fmt.Println("Gin server example configured (not running in this demo)")
	// router.Run(":8080") // เริ่ม server ที่ port 8080
}

// ==================== 18. GORM (Database ORM) ====================

// Database Models
type Customer struct {
	gorm.Model
	Name   string
	Email  string `gorm:"unique"`
	Orders []Order
}

type Order struct {
	gorm.Model
	CustomerID uint
	Product    string
	Quantity   int
	TotalPrice float64
	Customer   Customer
}

func gormExample() {
	// เชื่อมต่อ Database (SQLite)
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect database:", err)
		return
	}

	// Auto Migration (สร้างตาราง)
	db.AutoMigrate(&Customer{}, &Order{})

	// CREATE - สร้างข้อมูล
	customer := Customer{
		Name:  "John Smith",
		Email: "john.smith@example.com",
	}
	db.Create(&customer)

	// สร้าง Customer พร้อม Orders
	customerWithOrders := Customer{
		Name:  "Jane Doe",
		Email: "jane.doe@example.com",
		Orders: []Order{
			{Product: "Laptop", Quantity: 1, TotalPrice: 999.99},
			{Product: "Mouse", Quantity: 2, TotalPrice: 59.98},
		},
	}
	db.Create(&customerWithOrders)

	// READ - อ่านข้อมูล
	var customer1 Customer
	db.First(&customer1, 1) // หา customer ที่ ID = 1
	fmt.Printf("Customer: %+v\n", customer1)

	// Find by condition
	var customer2 Customer
	db.Where("email = ?", "john.smith@example.com").First(&customer2)

	// Get all
	var customers []Customer
	db.Find(&customers)
	fmt.Println("Total customers:", len(customers))

	// Preload (Load relations)
	var customerWithOrder Customer
	db.Preload("Orders").First(&customerWithOrder, 1)
	fmt.Printf("Customer with orders: %+v\n", customerWithOrder)

	// UPDATE - แก้ไขข้อมูล
	db.Model(&customer1).Update("Name", "John Updated")
	db.Model(&customer1).Updates(Customer{Name: "John Final", Email: "john.final@example.com"})

	// Save (Update all fields)
	customer1.Name = "John Save"
	db.Save(&customer1)

	// DELETE - ลบข้อมูล (Soft Delete)
	db.Delete(&customer1, 1)

	// Permanent Delete
	db.Unscoped().Delete(&customer1, 1)

	// QUERY - การค้นหาขั้นสูง
	var results []Customer

	// Where conditions
	db.Where("name LIKE ?", "%John%").Find(&results)
	db.Where("id > ?", 5).Find(&results)
	db.Where(map[string]interface{}{"name": "John", "email": "john@example.com"}).Find(&results)

	// Or condition
	db.Where("name = ?", "John").Or("email = ?", "jane@example.com").Find(&results)

	// Order, Limit, Offset
	db.Order("created_at desc").Limit(10).Offset(5).Find(&results)

	// Count
	var count int64
	db.Model(&Customer{}).Where("name LIKE ?", "%John%").Count(&count)
	fmt.Println("Count:", count)

	// Select specific fields
	db.Select("name", "email").Find(&results)

	// Group By & Having
	type Result struct {
		CustomerID uint
		Total      float64
	}
	var orderResults []Result
	db.Model(&Order{}).Select("customer_id, sum(total_price) as total").Group("customer_id").Having("total > ?", 100).Scan(&orderResults)

	// Joins
	db.Joins("JOIN orders ON orders.customer_id = customers.id").Find(&results)

	// Raw SQL
	db.Raw("SELECT * FROM customers WHERE name = ?", "John").Scan(&results)

	// Transaction
	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&Customer{Name: "Trans1", Email: "trans1@example.com"}).Error; err != nil {
			return err
		}
		if err := tx.Create(&Customer{Name: "Trans2", Email: "trans2@example.com"}).Error; err != nil {
			return err
		}
		return nil
	})

	// Hooks (Before/After actions)
	// func (u *Customer) BeforeCreate(tx *gorm.DB) error {
	//     // Do something before create
	//     return nil
	// }

	fmt.Println("GORM example completed")
}


func main() {
	fmt.Println("=== Go Programming Language Summary ===\n")

	fmt.Println("1. Variables Example:")
	variablesExample()

	fmt.Println("\n2. Data Types Example:")
	dataTypesExample()

	fmt.Println("\n3. Array & Slice Example:")
	arraySliceExample()

	fmt.Println("\n4. Map Example:")
	mapExample()

	fmt.Println("\n5. Operators Example:")
	operatorsExample()

	fmt.Println("\n6. Control Flow Example:")
	controlFlowExample()

	fmt.Println("\n7. Loops Example:")
	loopsExample()

	fmt.Println("\n8. Functions Example:")
	functionsExample()

	fmt.Println("\n9. Pointers Example:")
	pointersExample()

	fmt.Println("\n10. Structs Example:")
	structsExample()

	fmt.Println("\n11. Interfaces Example:")
	interfacesExample()

	fmt.Println("\n12. Error Handling Example:")
	errorHandlingExample()

	fmt.Println("\n13. Goroutines Example:")
	goroutinesExample()

	fmt.Println("\n14. Defer Example:")
	deferExample()

	fmt.Println("\n15. Recover Example:")
	recoverExample()

	fmt.Println("\n16. File I/O Example:")
	fileIOExample()

	fmt.Println("\n17. JSON Example:")
	jsonExample()

	fmt.Println("\n18. Gin Framework Example:")
	ginExample()

	fmt.Println("\n19. GORM Example:")
	gormExample()

	fmt.Println("\n=== Summary Complete ===")
	fmt.Println("This file contains comprehensive Go examples including:")
	fmt.Println("- Basic syntax and data types")
	fmt.Println("- Control flow and loops")
	fmt.Println("- Functions and methods")
	fmt.Println("- Structs and interfaces")
	fmt.Println("- Concurrency (Goroutines & Channels)")
	fmt.Println("- Error handling")
	fmt.Println("- GIN Framework for REST API")
	fmt.Println("- GORM for Database operations")
}