package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r *Rectangle) Area() float64 {
	return r.length * r.breadth
}

func PrintArea(s Shape) {
	fmt.Println("Area of s is : ", s.Area())
}

func main() {
	fmt.Println("Interface")
	c := Circle{4.3}

	PrintArea(&c)
	r := Rectangle{1.2, 3.4}
	PrintArea(&r)

}
