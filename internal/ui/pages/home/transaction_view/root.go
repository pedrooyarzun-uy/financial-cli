package transactionview

import (
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"
	"github.com/rivo/tview"
)

func NewTransactionsView(app *tview.Application, pages *tview.Pages) *tview.Flex {

	header, table, chart := NewTransactionsHeader(pages)
	viewGraphicBtn := components.NewButton("View Chart", func() {
		flex := tview.NewFlex().SetDirection(tview.FlexRow)
		flex.SetBorder(true).SetTitle("Chart transactions details")
		back := components.NewButton("Back", func() {
			pages.RemovePage("graphic_button")
		})

		flex.AddItem(chart, 0, 1, false)
		flex.AddItem(nil, 2, 0, true)
		flex.AddItem(back, 1, 40, true)

		pages.AddPage("graphic_button", flex, true, true)
	})

	bottomRow := tview.NewGrid().
		SetColumns(-1, -3, -1).
		SetRows(1)

	bottomRow.AddItem(tview.NewBox(), 0, 0, 1, 1, 0, 0, false)
	bottomRow.AddItem(viewGraphicBtn, 0, 1, 1, 1, 0, 0, false)
	bottomRow.AddItem(tview.NewBox(), 0, 2, 1, 1, 0, 0, false)

	root := tview.NewFlex().SetDirection(tview.FlexRow)
	root.SetBorder(true).SetTitle("View transactions")
	root.AddItem(header, 3, 0, true)
	root.AddItem(table, 0, 1, false)
	root.AddItem(bottomRow, 1, 0, false)

	return root
}
