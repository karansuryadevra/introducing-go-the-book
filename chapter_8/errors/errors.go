package main

import "errors"

func main() {
	err := errors.New("This is a newly created error message")
	panic(err)
}
