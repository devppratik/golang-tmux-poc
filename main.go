package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// The application.
var app = tview.NewApplication()
var pages = tview.NewPages()
var info = tview.NewTextView()

func main() {
	// Initial Slides
	tabSlides = append(tabSlides, *NewTab("kite", os.Getenv("SHELL")))
	tabSlides = append(tabSlides, *NewTab("bash", os.Getenv("SHELL")))

	// Input Methods and Handlers
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			nextSlide()
			return nil
		} else if event.Key() == tcell.KeyCtrlP {
			previousSlide()
			return nil
		} else if event.Key() == tcell.KeyCtrlA {
			tabSlide := *NewTab("bash", os.Getenv("SHELL"))
			tabSlides = append(tabSlides, tabSlide)
			pages.AddPage(strconv.Itoa(tabSlide.index), tabSlide.content, true, tabSlide.index == 0)
			fmt.Fprintf(info, `["%d"]%s[white][""]  `, tabSlide.index, fmt.Sprintf("%d %s", tabSlide.index+1, tabSlide.title))
			currentActivePage = tabSlide.index
			info.Highlight(strconv.Itoa(currentActivePage)).
				ScrollToHighlight()
		} else if event.Key() == tcell.KeyCtrlE {
			slideNum, _ := strconv.Atoi(info.GetHighlights()[0])
			Remove(info, slideNum)
			pages.RemovePage(strconv.Itoa(slideNum))
			previousSlide()
		}
		return event
	})

	// Init the app
	layout := initTerminalMux()
	// Start the application.
	if err := app.SetRoot(layout, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}
