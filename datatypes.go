package main

import "fmt"

func main() {

	//explicit declearation - we are defining the variable typre
	var name string = "chaitanya"
	var new string
	new = "sai"
	var number int
	number = 2
	number += 19

	//implicit declearation - no need to define the type
	var age2 = 25

	//most common used declearation
	new_number := 56

	//not defining
	var num int
	var b bool

	fmt.Println("My name is", name)
	fmt.Println("My Sister name is", new)
	fmt.Println("My age is", number)
	fmt.Printf("%T", age2) // "%T" gives the type of the variable
	fmt.Println(" ")
	fmt.Printf("%T", new_number) // "%T" gives the type of the variable
	fmt.Printf(" ")
	fmt.Println("Default value for int = ", num)
	fmt.Println("Default value for bool = ", b)

}
