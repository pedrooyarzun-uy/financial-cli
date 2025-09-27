package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewWarningModal(text string, pages *tview.Pages) *tview.Modal {
	modal := tview.NewModal().
		SetBackgroundColor(tcell.ColorYellow).
		SetTextColor(tcell.ColorBlack).
		SetText("\u26A0 " + text + " \u26A0").
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.RemovePage("modal")
		})

	modal.SetTitle("Warning")
	modal.SetTitleColor(tcell.ColorYellow)
	return modal

}
