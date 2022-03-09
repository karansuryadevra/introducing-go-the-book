package main

import "fmt"

func makeOddGenerator() func() int {
	i := 1
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}

func main() {

	oddGenerator := makeOddGenerator()
	fmt.Println(oddGenerator())
	fmt.Println(oddGenerator())
	fmt.Println(oddGenerator())

}
