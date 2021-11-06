package main

import "fmt"

//struct
type Point struct {
	x int
	y int
}

//embedded struct
type Circle struct {
	radius float32
	center *Point
}

func change_x(pt *Point) {
	pt.x = 100
}

func main() {
	var p1 Point = Point{4, 2}
	p2 := Point{4, 18}
	fmt.Println(p1.x, p2.x)
	fmt.Println(p1, p2)

	p3 := Point{x: 4} // other value is set to default value of that datatype
	fmt.Println(p3)

	// changing using refernces
	p4 := &Point{y: 5}
	fmt.Println(p4)
	change_x(p4)
	fmt.Println(p4)

	// Struct has a feature that for changing values
	// we dont need to deference the pointer like *p5.x = 100
	// we can directly write p5.x = 100
	p5 := &Point{44, 67}
	p5.x = 4213
	fmt.Println(p5)

	//embedded struct
	c1 := Circle{4.56, &p1}
	c2 := Circle{4.56, &Point{3, 5}}

	fmt.Println(c1, c1.center.x, c1.center.y)
	fmt.Println(c2)

}
