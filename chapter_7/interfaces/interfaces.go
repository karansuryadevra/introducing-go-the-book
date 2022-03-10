package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

type Shape interface { // Use this to abstract the properties of individual structs that have the same method (i.e. area())
	area() float64
}

type MultiShape struct { // Use an interface inside a struct/type definition. Allows us to add unlimited number of new shapes, as long as they have an area() function
	shapes []Shape
}

func (r Rectangle) area() float64 {
	return r.length * r.breadth
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func totalArea(shapes ...Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.area()
	}
	return total
}

func (m MultiShape) area() float64 {
	var area float64
	for _, v := range m.shapes { // m.shapes, not just 'm'
		area += v.area()
	}
	return area
}

func main() {
	c := Circle{2}
	r := Rectangle{5, 10}
	fmt.Println(c.area())
	fmt.Println(r.area())

	s := []Shape{c, r}
	fmt.Println(s) // create a slice that implements a shape type made of circles and rectangles that share area()
	fmt.Println(totalArea(&c, &r))

	multiShape := MultiShape{ // how to declare a type that has an interface inside it
		shapes: []Shape{
			Circle{2},
			Rectangle{5, 10},
		},
	}

	fmt.Println(multiShape.area())

}
