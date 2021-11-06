package main

import "fmt"

// operators: >, <, >=, <=, !=, ==
// || = OR
// && = AND
// ! = NOT

// for condition operators like +,-,*,/,%
//both the variables should be of same type
// like both int but not int and float
func main() {
	x := 5
	y := 6
	// val := x <= 5
	val := x == y

	// for string comparisions they ise ascii value of characters
	x1 := "a"
	y1 := "b"
	val1 := x1 < y1
	val2 := x1 == y1
	fmt.Printf("%t %t %t", val, val1, val2)
}
