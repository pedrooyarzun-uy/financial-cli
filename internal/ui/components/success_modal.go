package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewSuccessModal(text string, pages *tview.Pages) *tview.Modal {
	modal := tview.NewModal().
		SetBackgroundColor(tcell.ColorGreen).
		SetTextColor(tcell.ColorWhite).
		SetText(text).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.RemovePage("modal")
		})

	modal.SetTitle("Success")
	modal.SetTitleColor(tcell.ColorGreen)

	return modal
}
