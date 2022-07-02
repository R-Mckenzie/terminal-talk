package main

import (
	"net"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ui struct {
	app        *tview.Application
	msgDisplay *tview.TextView
	input      *tview.InputField
	flex       *tview.Flex
}

func initUI(conn net.Conn) ui {
	app := tview.NewApplication()

	tv := tview.NewTextView()
	tv.SetTitle("TERMINAL TALK").
		SetBorder(true)

	input := tview.NewInputField().SetLabel("You >> ")
	input.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			msg := input.GetText()
			conn.Write([]byte(msg))
			input.SetText("")
		case tcell.KeyEsc:
			app.Stop()
		}
	})

	flex := tview.NewFlex().
		SetDirection(tview.FlexColumnCSS).
		AddItem(tv, 0, 1, false).
		AddItem(input, 1, 1, true)

	app.SetRoot(flex, true)
	return ui{app, tv, input, flex}
}

func (ui *ui) printMessage(msg string) {
	ui.msgDisplay.Write([]byte(msg))
}
