package home

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewMain(app *tview.Application, pages *tview.Pages) *tview.Flex {

	main := tview.NewFlex()

	list := tview.NewList().
		AddItem("Add", "", 'a', func() {
			pages.SwitchToPage("add_page")
		}).
		AddItem("Automatic expense entry", "", 's', nil).
		AddItem("View stats", "", 'f', func() {
			pages.SwitchToPage("transactions_view")
		}).
		AddItem("Settings", "", 'g', nil).
		AddItem("Quit", "", 'q', func() {
			app.Stop()
		})

	list.SetShortcutColor(tcell.ColorOrange)

	main.AddItem(list, 0, 1, true)

	return main
}
