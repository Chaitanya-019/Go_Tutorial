package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func (s Student) getAge() int { // this function is only used by the student struct
	return s.age
}

func (s *Student) setAge(num int) { // this will actually modify the struct object
	s.age = num
}

func (s Student) getAvgGrade() float32 {
	sum := 0
	for _, val := range s.grades {
		sum += val
	}
	return float32(sum) / float32(len(s.grades))
}

func main() {
	s1 := Student{"Chaitu", []int{1, 2, 3, 4, 6}, 20}
	fmt.Println(s1.getAge())
	s1.setAge(19)
	fmt.Println(s1.age)
	fmt.Println("Average grades : ", s1.getAvgGrade())
}
