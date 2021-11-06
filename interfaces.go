package main

import "fmt"

type shape interface {
	area() float32
}

type Circle struct {
	radius float32
}

type Rect struct {
	length float32
	bredth float32
}

func (c *Circle) area() float32 {
	return 3.14 * (c.radius) * (c.radius)
}

func (r Rect) area() float32 {
	return r.length * r.bredth
}

func main() {
	c1 := Circle{4.5}
	r1 := Rect{4, 5}

	fmt.Println(c1.area())
	fmt.Println(r1.area())

	shapes := []shape{&c1, r1}

	for _, shap := range shapes {
		fmt.Println(shap.area())
	}
}
