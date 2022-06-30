package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

type server struct {
	clients []client
	msgChan chan string
}

func (s *server) run() {
	for message := range s.msgChan {
        log.Println(message)
		for _, c := range s.clients {
			c.send(message)
		}
	}
}

func main() {
	server := server{[]client{}, make(chan string)}
	go server.run()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		client := newClient(conn, server)
		server.clients = append(server.clients, *client)
		go client.listen()
	}
}
