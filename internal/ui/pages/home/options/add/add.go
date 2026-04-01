package add

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewAdd(app *tview.Application, pages *tview.Pages) *tview.Flex {
	pages.AddPage("add_transaction", NewTransaction(app, pages), true, true)
	pages.AddPage("add_account", NewAddAccount(app, pages), true, true)

	main := tview.NewFlex()

	list := tview.NewList().
		AddItem("Add expense", "", 'a', func() {
			pages.SwitchToPage("add_transaction")
		}).
		AddItem("Add account", "", 's', func() {
			pages.SwitchToPage("add_account")
		})

	list.SetShortcutColor(tcell.ColorOrange)
	main.AddItem(list, 0, 1, true)

	return main
}
