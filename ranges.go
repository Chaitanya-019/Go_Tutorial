package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 53, 2, 3, 2, 1}

	for i, element := range a { // similar to enumerate in python, i = index and element = a[i], if i or element is not used repalce it with ' _ '
		fmt.Println(i, element)
	}

	//printng only dublicates numbers
	fmt.Println("Printing duplicates")
	for i, element1 := range a {
		for j := i + 1; j < len(a); j++ {
			// fmt.Println(i, j, element1, a[j])
			if a[j] == element1 {
				fmt.Println(element1)
			}
		}
	}

}
