package main

import "fmt"

func main() {
	var x [5]int = [5]int{1, 2, 3, 4, 5} // array

	//slice
	var s []int = x[:]   // full x
	var s1 []int = x[1:] // from one index to end of x
	var s2 []int = x[:3] // from start to 2nd index

	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(len(s1)) // length
	fmt.Println(len(s2))
	fmt.Println(s2[:cap(s2)]) // capacity = what can be the capacity of slice as it from x , it has capacity 5

	//making slices
	// here the slice cannot be extended but needed to create new slice with appending new value
	// if we need same variable we can do below implementation
	var a []int = []int{3, 2, 4, 5, 6}
	a = append(a, 53)
	fmt.Println(a)

	//way - 2
	b := make([]int, 5) // an array of size 5
	fmt.Println((b))

}
