package home

import (
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/pages/home/options"
	"github.com/rivo/tview"
)

func NewHomePage(app *tview.Application, pages *tview.Pages) *tview.Grid {

	sidebar := NewSidebar(app, pages)

	main := NewMain(app, pages)
	pages.AddPage("transaction", options.NewTransaction(app, pages), true, true)

	//Seteamos focus a main
	app.SetFocus(main)

	header := NewHeader()

	grid := tview.NewGrid().
		SetRows(2, 2).
		SetColumns(30, 0).
		SetBorders(true)

	grid.AddItem(header, 0, 0, 1, 2, 0, 0, false)
	grid.AddItem(main, 1, 0, 2, 1, 0, 100, true)
	grid.AddItem(sidebar, 1, 1, 2, 1, 0, 100, false)

	return grid
}
