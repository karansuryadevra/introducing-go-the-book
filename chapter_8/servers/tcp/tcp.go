package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server(ln net.Listener) {

	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	var msg string
	err := gob.NewDecoder(c).Decode(&msg) // see what happens
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received the following message: ", msg)
	}
	c.Close()
}

func client() {
	c, err := net.Dial("tcp", "127.0.0.1:9689")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "Hello, World!"
	fmt.Println("Sending the following message: ", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Close()
}
func main() {
	ln, err := net.Listen("tcp", ":9689")
	if err != nil {
		panic(err)
	}
	go server(ln)
	go client()
	var input string
	fmt.Scanln(&input)
}
