package main

import "fmt"

// & = pointer
// * = dereference

// as this is global this can be changed any where in the code
var str1 string = "saikala"

// taking the pointers
func change_value(str *string) {
	*str = "Good boy"
	str1 = "Good girl"
}

func main() {
	x := 7
	fmt.Println("Reference : ", &x, " Value: ", x) // this gives reference or adress it is stored

	y := &x // now y is the refernce of x address : NOW IF I CHANGE Y then x also changes because it is the address where x is pointing

	fmt.Println("Before changing in address : ", x, y)

	*y = 8 // changing the value at y = addaress

	fmt.Println("After changing in address : ", x, y)

	str := "chaitanya"
	// this is giving the pointer
	change_value(&str)
	fmt.Println(str)

	fmt.Println(str1)

}
