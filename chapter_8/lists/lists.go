package main

import (
	"container/list"
	"fmt"
)

func main() {
	var x list.List
	x.PushBack(1) // PushBack appends a new value to the list
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() { // remember the for loop syntax: for initialization; condition; post iteration{}
		fmt.Println(e.Value.(int))
	}

	fmt.Println(x)
}
