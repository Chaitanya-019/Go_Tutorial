package main

import "fmt"

func main() {
	var arr [5]int
	arr[2] = 3
	fmt.Println(arr)

	var arr2d [5][4]int
	fmt.Println(arr2d)

	arr3 := []int{-1}
	fmt.Println(arr3)

	arr4 := [6]int{-1}
	fmt.Println(len(arr4))

	arr5 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr5, len(arr5), len(arr5[0]))

	sum := 0
	for x := 0; x < len(arr4); x++ {
		sum += arr4[x]
	}
	fmt.Println("sum = ", sum)
}
