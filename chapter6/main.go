package main

import "fmt"

func average(xs []float64) float64{
	total := 0.0

	for _, val := range xs{
		total +=val
	}

	return total / float64(len(xs))
}
/*
	"add" is called a variadic function. Basically, that allows the caller function to pass multiple
	values of a particular type to it. 
	Therefore the main can pass multiple integers to the add function
*/
func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}

func makeEvenGenerator() func() uint{
	i := uint(0)
	return func() (ret uint){
		ret = i
		i += 2
		return
	}
}

func main(){
	xs := []float64{98,93,77,82,83}
	fmt.Println(average(xs))
	fmt.Println(add(1,2,3,4,5))

	// declaring a function within a function :-
	x := 0
	increment := func() int{
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	// This whole block above of a function within a function and the local vars is called "Closure"

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

}