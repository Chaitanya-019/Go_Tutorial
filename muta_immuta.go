package main

import "fmt"

// mutable datatype = when a new element is assigned from this present element of same datatype then, change in any one of them effects the  both
// example: slices, maps

// immutable  datatype= when a new element is assigned from this present element of same datatype then, change in any one of them, doesnt effect other
// example: int, bool float,array ...

func main() {
	var x int = 4
	y := x
	y = 4
	fmt.Println(x, y)

	var a [3]int = [3]int{4, 2, 4}
	b := a
	b[2] = 423
	fmt.Println(a, b)

	var x1 []int = []int{3, 52, 5}
	y1 := x1
	y1[0] = 100
	fmt.Println(x1, y1)

	var x2 map[string]int = map[string]int{"apple": 30}
	y2 := x2
	y2["wow"] = 4
	fmt.Println(x2, y2)

}
