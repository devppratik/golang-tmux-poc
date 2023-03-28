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
var inputBuffer []rune

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
		} else if event.Key() == tcell.KeyRune {
			inputBuffer = append(inputBuffer, event.Rune())
			if string(inputBuffer) == "exit\n" {
				app.Stop()
				return nil
			}
		} else if event.Key() == tcell.KeyBackspace || event.Key() == tcell.KeyBackspace2 {
			if len(inputBuffer) > 0 {
				inputBuffer = inputBuffer[:len(inputBuffer)-1]
			}
		} else if event.Key() == tcell.KeyEnter {
			if string(inputBuffer) == "exit" {
				// Clear input buffer
				inputBuffer = []rune{}
				app.Stop()
				return nil
			}
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
