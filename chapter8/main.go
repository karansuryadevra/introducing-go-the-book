package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	bs, err := ioutil.ReadFile("test.txt")
	
	if err != nil {
		return
	}

	fmt.Println(string(bs))

}
