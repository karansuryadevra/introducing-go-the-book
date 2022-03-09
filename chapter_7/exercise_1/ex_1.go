package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float32
}

type Rectangle struct {
	length  float32
	breadth float32
}

type Shape interface {
	area() float32
	perimeter() float32
}

func (c Circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

func (c Circle) perimeter() float32 {
	return 2 * math.Pi * c.radius
}

func (r Rectangle) perimeter() float32 {
	return 2 * (r.length + r.breadth)
}

func totaler(shapes ...Shape) (float32, float32) {
	area := float32(0)
	perimeter := float32(0)
	for _, v := range shapes {
		area += v.area()
		perimeter += v.perimeter()
	}
	return area, perimeter
}

func main() {
	shape := []Shape{Circle{2}, Rectangle{5, 4}}
	fmt.Println(totaler(shape...))
}
