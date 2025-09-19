package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewHomePage(app *tview.Application, pages *tview.Pages) *tview.Flex {
	box := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText("Welcome Back!").SetTextColor(tcell.ColorGreen), 2, 0, false)
	return box
}
