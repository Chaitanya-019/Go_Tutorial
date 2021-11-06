package main

import (
	"fmt"
)

func main() {
	//order is not matained
	// similar to unordered_map in c++
	var mp map[string]int = map[string]int{
		"apple": 5,
		"car":   43,
		"lambo": 42,
	} //mp[string] = int
	mp["apple"] = 0   //modifing
	delete(mp, "car") // deleting
	fmt.Println(mp)

	mp1 := make(map[string]int) //an empty map
	mp1["chaitanya"] = 100      // adding new value
	fmt.Println(mp1)

	val, ok := mp1["sai"] // val is the value of key if exist and default if not, ok is true or false to check if it is there or not
	fmt.Println(val, ok)

	// for loop
	for key, val := range mp {
		fmt.Println(key, val)
	}

}
