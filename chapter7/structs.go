package main 

import ("fmt"; "math")

func (c *Circle) area() float64 { 
	/* 
		area() is called a receiver and allows you to use "."
		to access the function from the caller
	*/
	return math.Pi * c.r*c.r
}

//declare a struct
type Circle struct {
	x, y, r float64
}

func main(){ 

	// initialze a Circle struct
	c := Circle{x: 0, y: 5, r: 5 }
	// or
	d := Circle{0, 5, 5}
	// or (less commonly used) which will set all values to 0 and return a pointer to Circle i.e. &{0,0,0}
	e := new(Circle)

	fmt.Println(c,d,e)

	fmt.Println(e.area())	
}