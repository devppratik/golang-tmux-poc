package main

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"
)

func Remove(info *tview.TextView, s int) {
	tabSlides = append(tabSlides[:s], tabSlides[s+1:]...)
	regionIds = append(regionIds[:s], regionIds[s+1:]...)
	totalPageCount = len(tabSlides)
	info.Clear()
	for index, tabSlide := range tabSlides {
		oldIndex := tabSlide.index
		tabSlide.index = index
		fmt.Fprintf(info, `["%d"]%s[white][""]  `, oldIndex, fmt.Sprintf("%d %s", tabSlide.index+1, tabSlide.title))
	}
}

func previousSlide() {
	currentActivePage = (currentActivePage - 1 + len(tabSlides)) % len(tabSlides)
	info.Highlight(strconv.Itoa(regionIds[currentActivePage])).
		ScrollToHighlight()
}
func nextSlide() {
	currentActivePage = (currentActivePage + 1) % len(tabSlides)
	info.Highlight(strconv.Itoa(regionIds[currentActivePage])).
		ScrollToHighlight()
}

func initTerminalMux() *tview.Flex {
	info.
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			pages.SwitchToPage(added[0])
		})

	for _, tabSlide := range tabSlides {
		pages.AddPage(strconv.Itoa(tabSlide.index), tabSlide.content, true, tabSlide.index == 0)
		fmt.Fprintf(info, `["%d"]%s[white][""]  `, tabSlide.index, fmt.Sprintf("%d %s", tabSlide.index+1, tabSlide.title))
	}
	info.Highlight("0")
	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false)
}
