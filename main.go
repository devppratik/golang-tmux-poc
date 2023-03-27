package main

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// The application.
var app = tview.NewApplication()
var pages = tview.NewPages()
var info = tview.NewTextView()

func main() {
	// App Shorcuts Implementation
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			nextSlide()
			return nil
		} else if event.Key() == tcell.KeyCtrlP {
			previousSlide()
			return nil
		} else if event.Key() == tcell.KeyCtrlA {
			addSlide()
		} else if event.Key() == tcell.KeyCtrlE {
			slideNum, _ := strconv.Atoi(info.GetHighlights()[0])
			removeSlide(slideNum)
		}
		return event
	})

	// Get the initial Config
	layout := initTerminalMux()
	// Start the application.
	if err := app.SetRoot(layout, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}
