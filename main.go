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

// Starting point for the presentation.
func main() {
	// The presentation slides.
	tabSlides = append(tabSlides, *NewTab("kite", os.Getenv("SHELL")))
	tabSlides = append(tabSlides, *NewTab("bash", os.Getenv("SHELL")))

	pages := tview.NewPages()

	// The bottom row has some info on where we are.
	info := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			pages.SwitchToPage(added[0])
		})

	// Create the pages for all slides.
	previousSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide - 1 + len(tabSlides)) % len(tabSlides)
		info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}
	nextSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide + 1) % len(tabSlides)
		info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}
	for _, tabSlide := range tabSlides {
		pages.AddPage(strconv.Itoa(tabSlide.index), tabSlide.content, true, tabSlide.index == 0)
		fmt.Fprintf(info, `["%d"]%s[white][""]  `, tabSlide.index, fmt.Sprintf("%d %s", tabSlide.index+1, tabSlide.title))
	}
	info.Highlight("0")

	// Create the main layout.
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false)

	// Shortcuts to navigate the slides.
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
			info.Highlight(strconv.Itoa(tabSlide.index)).
				ScrollToHighlight()
		} else if event.Key() == tcell.KeyCtrlE {
			slideNum, _ := strconv.Atoi(info.GetHighlights()[0])
			Remove(info, slideNum)
			pages.RemovePage(strconv.Itoa(slideNum))
			previousSlide()
		}
		return event
	})

	// Start the application.
	if err := app.SetRoot(layout, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}
