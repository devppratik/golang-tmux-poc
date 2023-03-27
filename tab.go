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
var regionIds []int
var currentActivePage int = 0
var totalPageCount int = -1

func NewTab(name string, command string) *tab {
	totalPageCount += 1
	regionIds = append(regionIds, totalPageCount)
	return &tab{
		index:   totalPageCount,
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
