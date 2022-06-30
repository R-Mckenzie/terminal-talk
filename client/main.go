package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	name := os.Args[1]
	conn.Write([]byte(fmt.Sprint(name + "\n")))

	go listenTCP(conn)
	readInput(conn)
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
