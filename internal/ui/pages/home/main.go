package home

import "github.com/rivo/tview"

func NewMain(app *tview.Application) *tview.Flex {
	main := tview.NewFlex()

	list := tview.NewList().
		AddItem("Add expense", "", 'a', nil).
		AddItem("Automatic expense entry", "", 's', nil).
		AddItem("Add account", "", 'd', nil).
		AddItem("View stats", "", 'f', nil).
		AddItem("Settings", "", 'g', nil).
		AddItem("Quit", "", 'q', func() {
			app.Stop()
		})

	main.AddItem(list, 0, 1, true)

	return main
}
