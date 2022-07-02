package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type App struct {
	incoming chan string
	outgoing chan string
	conn     net.Conn
	ui       ui
}

func appInit() App {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	// Sends server the user nickname
	name := os.Args[1]
	conn.Write([]byte(fmt.Sprint(name + "\n")))
	outgoing := make(chan string)
	app := App{make(chan string), outgoing, conn, initUI(outgoing)}
	return app
}

func (a *App) close() {
	//cleanup
	a.conn.Close()
}

func main() {
	app := appInit()
	defer app.close()

	go app.listenTCP()
	go app.sendMessages()

	if err := app.run(); err != nil {
		panic(err)
	}
}

func (a *App) run() error {
	return a.ui.app.Run()
}

func (a *App) sendMessages() {
	for msg := range a.outgoing {
		a.ui.printMessage(msg)
		a.conn.Write([]byte(msg + "\n"))
	}
}

func (a *App) printIncoming(msg string) {
	a.ui.printMessage(msg)
	a.ui.app.Draw()
}

// Listens to incoming TCP connections and sends them to a channel
func (a *App) listenTCP() {
	reader := bufio.NewReader(a.conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Lost connection to the server")
			os.Exit(1)
		}
		newMsg := strings.TrimSpace(string(message))
		a.printIncoming(newMsg)
	}
}
