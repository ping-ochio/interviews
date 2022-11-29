package main

import "fmt"

const pi = 3.14159

type Figurer interface {
	Area() (float32, string)
	Perimeter() (float32, string)
}

type Circle struct {
	radio float32
}

type Rectangle struct {
	high  float32
	width float32
}

type Figus struct {
	cir Circle
	rec Rectangle
}

func (c *Circle) Area() (float32, string) {

	return (pi * (c.radio * c.radio)), "Circle"

}

func (q *Rectangle) Area() (float32, string) {

	return (q.high * q.width), "Rectangle"

}

func (c *Circle) Perimeter() (float32, string) {
	return (2 * pi * c.radio), "Circle"

}

func (q *Rectangle) Perimeter() (float32, string) {
	return (2*q.high + 2*q.width), "Rectangle"
}

func Calcarea(f Figurer) {
	area, figure := f.Area()
	fmt.Printf("The %s area is: %v [u^2]\n", figure, area)
}

func Calcper(f Figurer, n string) {
	perimeter, figure := f.Perimeter()
	fmt.Printf("The %s named %s perimeter is: %v [u]\n", figure, n, perimeter)
}

func main() {

	cir1 := Circle{4}
	cir2 := Circle{3}
	rect1 := Rectangle{6, 4}
	rect2 := Rectangle{4, 4}

	//lala := []Figus{{Circle{5}, Rectangle{5, 6}}, {Circle{2}, Rectangle{8, 9}}}
	//fmt.Println(lala)
	Calcarea(&cir1)
	Calcarea(&cir2)
	Calcarea(&rect1)
	Calcarea(&rect2)

	Calcper(&cir1, "cir1")
	Calcper(&cir2, "cir2")
	Calcper(&rect1, "rect1")
	Calcper(&rect2, "rect2")

}
