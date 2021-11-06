package main

import (
	"fmt"
	// "math" // this has alll the math functions
)

// for arthematic operators like +,-,*,/,%
//both the variables should be of same type
// like both int but not int and float
func main() {

	// float(4) = takes 4 as a float val
	// int(3.2652) = 3
	num1 := 9
	num2 := 4
	answer := num1 / num2 // answer will be same datatype as num1 and num2
	fmt.Printf("%d", answer)
}
