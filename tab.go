package main

import (
	"os"
	"os/exec"

	"git.sr.ht/~rockorager/tterm"
	"github.com/rivo/tview"
)

type tab struct {
	index   int
	title   string
	content tview.Primitive
}

var tabSlides []tab

func NewTab(name string) *tab {
	var index int
	if len(tabSlides) == 0 {
		index = 0
	} else {
		index = len(tabSlides)
	}
	return &tab{
		index:   index,
		title:   name,
		content: NewTabSlide(),
	}
}

func NewTabSlide() (content tview.Primitive) {
	cmd := exec.Command(os.Getenv("SHELL"))
	term := tterm.NewTerminal(cmd)
	term.SetBorder(true)
	return term
}
