package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type App struct {
	incoming chan string
	outgoing chan string
	conn     net.Conn
}

func start() App {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	// Sends server the user nickname
	name := os.Args[1]
	conn.Write([]byte(fmt.Sprint(name + "\n")))
	return App{make(chan string), make(chan string), conn}
}

func main() {
	app := start()
	defer app.conn.Close()

	ui := initUI(app.conn)
	go app.listenTCP(ui)
	if err := ui.app.Run(); err != nil {
		panic(err)
	}

}

// Reads input from the user
func (a *App) readInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input: %v", err)
			continue
		}
		a.conn.Write([]byte(input))
	}
}

// Listens to incoming TCP connections and sends them to a channel
func (a *App) listenTCP(ui ui) {
	reader := bufio.NewReader(a.conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Lost connection to the server")
			os.Exit(1)
		}
		out := strings.TrimSpace(string(message))
		a.incoming <- out
		ui.printMessage(out)
	}
}
