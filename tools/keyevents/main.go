package main

import (
	"fmt"

	tcell "github.com/gdamore/tcell/v2"
	"github.com/yellowsink/gord/tview"
)

func main() {
	input := tview.NewTextView()
	input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		fmt.Fprintln(input, event.Key(), event.Modifiers(), event.Rune())

		return event
	})
	tview.NewApplication().SetRoot(input, true).Run()
}
