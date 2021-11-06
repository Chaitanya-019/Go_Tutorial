package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	//taking input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type anything: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("You typed: %q \n", input)

	//converting string to int and many more
	scanner1 := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type your birth year : ")
	scanner1.Scan()
	input1, _ := strconv.ParseInt(scanner1.Text(), 10, 64)
	fmt.Printf("You will be %d years old by end of 2020", 2020-input1)

}
