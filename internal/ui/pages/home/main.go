package home

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewMain(app *tview.Application, pages *tview.Pages) *tview.Flex {

	main := tview.NewFlex()

	list := tview.NewList().
		AddItem("Add expense", "", 'a', func() {
			pages.SwitchToPage("transaction")
		}).
		AddItem("Automatic expense entry", "", 's', nil).
		AddItem("Add account", "", 'd', nil).
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
