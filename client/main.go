package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	Host = "localhost"
	Port = "8080"
	Type = "tcp"
)

func main() {
	fmt.Printf("Starting %s connection on %s:%s\n", Type, Host, Port)

	name := os.Args[1]
	conn, err := net.Dial(Type, Host+":"+Port)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	conn.Write([]byte(fmt.Sprint(name + "\n")))

	go listenTCP(conn)
	readInput(conn)
	return
}

func readInput(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input: %v", err)
			continue
		}
		conn.Write([]byte(input))
	}
}

func listenTCP(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Lost connection to the server")
			os.Exit(1)
		}
		out := strings.TrimSpace(string(message))
		log.Println(out)
	}
}
