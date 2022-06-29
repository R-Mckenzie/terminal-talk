package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	Host = "localhost"
	Port = "8080"
	Type = "tcp"
)

func main() {
	fmt.Printf("Starting %s connection on %s:%s\n", Type, Host, Port)

	ln, err := net.Listen(Type, Host+":"+Port)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error connecting: ", err.Error())
			return
		}
		fmt.Println("Client connected.")
		fmt.Println("Client " + conn.RemoteAddr().String() + " connected.")
	}
}

func handleConnection(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}
	log.Println("Client message: ", string(buffer[:len(buffer)-1]))
	conn.Write(buffer)
	handleConnection(conn)
}
