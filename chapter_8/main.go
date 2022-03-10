package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test", "te"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println(strings.HasSuffix("test", "st"))
	fmt.Println(strings.Index("test", "e"))
	fmt.Println(strings.Join([]string{"test", "e"}, ","))
	fmt.Println(strings.Repeat("test", 2))
	fmt.Println(strings.Split("test", "e"))
	fmt.Println(strings.ToUpper("test"))

	arr := []byte("String")
	fmt.Println(arr)
	str := string(arr)
	fmt.Println(str)

	var buf bytes.Buffer
	buf.Write(arr)
}
