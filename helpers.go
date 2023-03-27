package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func Remove(info *tview.TextView, s int) {
	tabSlides = append(tabSlides[:s], tabSlides[s+1:]...)
	info.Clear()
	for index, tabSlide := range tabSlides {
		tabSlide.index = index
		fmt.Fprintf(info, `["%d"]%s[white][""]  `, tabSlide.index, fmt.Sprintf("%d %s", tabSlide.index+1, tabSlide.title))
	}
}
