package main

import "fmt"

func main() {
	// for loop
	// style 1
	x := 4
	for x < 5 {
		fmt.Println("style 1")
		x++
	}

	//style 2
	for x := 0; x < 1; x += 2 {
		fmt.Println("style 2")

	}

	//break and continue
	for x := 0; x < 5; x += 1 {
		if x == 1 {
			fmt.Println(x)
			continue
		}
		if x == 2 {
			break
		}
	}
}
