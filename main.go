package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"git.sr.ht/~rockorager/tterm"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Slide func(nextSlide func()) (title string, content tview.Primitive)

// The application.
var app = tview.NewApplication()

// Starting point for the presentation.
func main() {

	// The presentation slides.
	slides := []Slide{Bash}

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
		title, primitive := slide(nextSlide)
		pages.AddPage(strconv.Itoa(index), primitive, true, index == 0)
		fmt.Fprintf(info, `%d ["%d"][darkcyan]%s[white][""]  `, index+1, index, title)
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
		} else if event.Key() == tcell.KeyCtrlM {
			previousSlide()
			return nil
		} else if event.Key() == tcell.KeyCtrlA {
			slides = append(slides, Python)
			title, primitive := Python(nextSlide)
			index := len(slides) - 1
			pages.AddPage(strconv.Itoa(index), primitive, true, false)
			fmt.Fprintf(info, `%d ["%d"][darkcyan]%s[white][""]  `, index+1, index, title)
		}
		return event
	})

	// Start the application.
	if err := app.SetRoot(layout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func Bash(nextSlide func()) (title string, content tview.Primitive) {
	cmd := exec.Command(os.Getenv("SHELL"))
	term := tterm.NewTerminal(cmd)

	term.SetBorder(true)
	term.SetBorderPadding(1, 1, 1, 1)
	term.SetTitle(" Welcome to bash")
	term.SetTitleColor(tcell.ColorBlue)
	return "bash", term
}

func Python(nextSlide func()) (title string, content tview.Primitive) {
	containerName := "python"
	cmd := exec.Command("python")
	term := tterm.NewTerminal(cmd)

	term.SetTitle(containerName)
	term.SetTitleColor(tcell.ColorBlue)
	return containerName, term
}
