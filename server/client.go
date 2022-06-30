package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type client struct {
	conn    net.Conn
	reader  *bufio.Reader
	msgChan chan message
	name    string
}

func newClient(conn net.Conn, server server) *client {
	reader := bufio.NewReader(conn)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	c := &client{conn, reader, server.msgChan, name}
	return c
}

func (c *client) listen() {
	for {
		msg, err := c.reader.ReadString('\n')
		if err != nil {
			log.Printf("%s has disconnected", c.name)
			return
		}
		msg = strings.TrimSpace(msg)
		msg = c.name + ": " + msg
		c.msgChan <- message{msg, *c}
	}
}

func (c *client) send(message string) {
	c.conn.Write([]byte(message + "\n"))
}
