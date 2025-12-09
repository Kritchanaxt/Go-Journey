package main

import "fmt"

func main() {
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