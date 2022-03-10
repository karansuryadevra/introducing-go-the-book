package main

import "fmt"

func averageFunc(xs []float64) float64 {
	total := 0.0

	for _, v := range xs {
		total += v
	}

	return total / float64(len(xs))
}

func multipleValuesReturner() (int, int) {
	return 1, 2
}

func variadicFunction(args ...int) int { // takes in a variable number of params. Just like fmt.Println(args...interface{})(n int, err error)
	total := 0

	for _, v := range args {
		total += v
	}
	return total
}

func makeEvenNumberGenerator() func() uint { //function that returns another function
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func recursiveFactorial(fact uint) uint {
	if fact == 0 {
		return 1
	}

	return fact * recursiveFactorial(fact-1)
}

func firstFunc() {
	fmt.Println("First")
}

func secondFunc() {
	fmt.Println("Second")
}

func pointerFunction(xPtr *int) {
	*xPtr = 0
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(averageFunc(xs))

	first, second := multipleValuesReturner()
	fmt.Println("Mutliple Values can be returned like this: ", first, " ", second)

	fmt.Println(variadicFunction(1, 2, 3, 4, 5, 6, 7, 8, 9))
	sliceOfNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(variadicFunction(sliceOfNumbers...)) // Note the syntax and how the same function is also allowing a slice: The beauty of a variadic function

	// closure begins here
	iterator := 0
	incrementer := func() int { // Called a local function. Has access to local variables
		iterator++
		return iterator
	}
	// closure ends here
	fmt.Println(incrementer())
	fmt.Println(incrementer())

	nextEvenNumber := makeEvenNumberGenerator() // use a function that returns a function. Allows a local variable value to persist between function calls
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())

	fmt.Println(recursiveFactorial(3))

	pointerVar := 22
	pointerFunction(&pointerVar) // &pointerVar returns a *int since pointerVar is 22 which is an int
	fmt.Println(pointerVar)      // will print 0 insteda of 22 even though the variable value was changed in the function

	newPointerVar := new(int)      // creates a *int
	pointerFunction(newPointerVar) // dont need to use &newPointerVar because that is the same as what newPointerVar is i.e. *int
	fmt.Println(*newPointerVar)

	defer firstFunc() // will execute after panic below (LIFO since there are 2 defer functions)
	secondFunc()

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC!!!")

}
