package main

import (
	"encoding/go"
	"fmt"
	"net"
)

func server(){
	//Listen on port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil{
		fmt.Println(err)
		return
	}
	for {
		//accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//handle the connection
		go handleServerConnection(c)		
	}
}

func handleServerConnection(c net.Conn){
	//receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	c.Close()
}

func client(){
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1")
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Close()
}

func main(){
	go server()
	go clint()

	var input string
	fmt.Scanln(&input)
}