package main

import "fmt"

func test(x uint8) {
	fmt.Println(x)
}

// giving every variable datatype
func add(x int, y int) {
	fmt.Println("sum - ", x+y)
}

// all datatypes are same
func sub(x, y int, z string) {
	fmt.Println("diff - ", x-y)
	fmt.Println(z)
}

// returning only 1
func multi(x, y int) int {
	return x * y
}

//multi return
func divide(x, y int) (int, int) {
	return x / y, x % y
}

// other way
func divide1(x, y int) (z1 int, z2 int) {
	defer fmt.Println("This is defer and runs only before return statement")
	z1 = x / y
	z2 = x % y
	fmt.Println("This is before return which executes before defer")
	return
}

// giving function as a parameter to a function
func new_func(my_func func(int) int) {
	fmt.Println(my_func(422))
}

// returning func
func returnfunc(x string) func() { // here returning does take parameters and dont return anything
	return func() {
		fmt.Println(x)
	}
}

func main() {
	add(5, 7)
	sub(5, 2, "wow")
	fmt.Println(multi(3, 2))
	ans1, ans2 := divide(4, 6)
	fmt.Println(ans1, ans2)
	fmt.Println(divide1(4, 2))

	// referencing function as pointer
	x := test // this 2 lines is same as test()
	x(5)

	// inside function
	test1 := func(x int) {
		fmt.Println("x = ", x)
	}
	test1(4)

	// value directly from a func
	test2 := func(x int) int {
		return -1 * x
	}(4)
	fmt.Println(test2)

	test3 := func(c int) int {
		return -1 * c
	}

	new_func(test3)

	returnfunc("hello")()
	z := returnfunc("chaitanya")
	z()

}
