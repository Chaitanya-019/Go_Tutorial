package main

import "fmt"

func main() {
	name := "sai"

	if name == "asi" {
		fmt.Println("Name is asi")
	} else if name == "sai" {
		fmt.Println("Name is sai")

	} else {
		fmt.Println("Name is not asi and sai")
	}

	//switch statement
	ans := 10
	switch ans {
	case 1, -1: // this is like 1 or -1
		fmt.Println("ans = 1 or -1")
	case 2:
		fmt.Println("ans = ", 2)
	default:
		fmt.Println("ans !=1 and ans=!=2")
	}
	switch {
	case ans >= 0:
		fmt.Println("ans>0")
	}

}
