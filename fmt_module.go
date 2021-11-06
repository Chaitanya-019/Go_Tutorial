package main

import "fmt"

//FMT cheat sheet
func main() {
	fmt.Printf("Hello %T %v", 10, 10)
	var s string = fmt.Sprintf("Hey there")
	var s2 string = fmt.Sprintf("Hello %T %v", 10, 10)
	fmt.Println(s2)
	fmt.Println(" ")
	fmt.Println(s)
	fmt.Printf("Hello %T %v %t", 10, 10, true)
	fmt.Println()
	fmt.Printf("Floating %T %v %e", 10.2665, 10.2665, 10.2665)
	fmt.Println(" ")

	//width and precision
	fmt.Printf("Name : %20q \n", "chaitanya")
	fmt.Printf("Name : %-20q is cool \n", "chaitanya")
	fmt.Printf("float : %.2f \n", 165.654984)

	//padding
	fmt.Printf("Number: %05d", 301)
	fmt.Println(" ")
}
