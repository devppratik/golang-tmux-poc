package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"git.sr.ht/~rockorager/tterm"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Slide func(name string, nextSlide func()) (title string, content tview.Primitive)

// The application.
var app = tview.NewApplication()

// Starting point for the presentation.
func main() {
	// The presentation slides.
	slides := []Slide{
		NewSlide,
		NewSlide,
	}

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
		slide = (slide - 1 + len(slides)) % len(slides)
		info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}
	nextSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide + 1) % len(slides)
		info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}
	for index, slide := range slides {
		title, primitive := slide("kite", nextSlide)
		pages.AddPage(strconv.Itoa(index), primitive, true, index == 0)
		fmt.Fprintf(info, `["%d"]%s[white][""]  `, index, fmt.Sprintf("%d %s", index+1, title))
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
			slides = append(slides, NewSlide)
			title, primitive := NewSlide("ocm-container", nextSlide)
			index := len(slides) - 1
			pages.AddPage(strconv.Itoa(index), primitive, true, false)
			fmt.Fprintf(info, `["%d"]%s[white][""]  `, index, fmt.Sprintf("%d %s", index+1, title))
		} else if event.Key() == tcell.KeyCtrlE {
			slideNum, _ := strconv.Atoi(info.GetHighlights()[0])
			slides = Remove(slides, slideNum)
			pages.RemovePage(strconv.Itoa(slideNum))
			currText := info.GetText(false)
			remText := info.GetRegionText(info.GetHighlights()[0])
			res := strings.ReplaceAll(currText, remText, "")
			previousSlide()
			info.Clear()
			fmt.Fprint(info, res)

		}
		return event
	})

	// Start the application.
	if err := app.SetRoot(layout, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}

func NewSlide(name string, nextSlide func()) (title string, content tview.Primitive) {
	cmd := exec.Command(os.Getenv("SHELL"))
	term := tterm.NewTerminal(cmd)
	term.SetBorder(true)
	term.SetTitle(" Welcome to kite ")
	term.SetTitleColor(tcell.ColorBlue)
	return name, term
}

func Remove(slice []Slide, s int) []Slide {
	return append(slice[:s], slice[s+1:]...)
}
