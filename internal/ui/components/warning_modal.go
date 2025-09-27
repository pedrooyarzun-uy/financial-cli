package components

import "github.com/rivo/tview"

func NewWarningModal(text string, pages *tview.Pages) *tview.Modal {
	modal := tview.NewModal().
		SetText(text).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.RemovePage("modal")
		})

	return modal

}
