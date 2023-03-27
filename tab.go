package main

import (
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

func NewTab(name string, command string) *tab {
	index := len(tabSlides)
	if len(tabSlides) == 0 {
		index = 0
	}
	return &tab{
		index:   index,
		title:   name,
		content: NewTabSlide(command),
	}
}

func NewTabSlide(command string) (content tview.Primitive) {
	cmd := exec.Command(command)
	term := tterm.NewTerminal(cmd)
	term.SetBorder(true)
	return term
}
