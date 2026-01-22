package transactionview

import "github.com/rivo/tview"

func NewTransactionsView(app *tview.Application, pages *tview.Pages) *tview.Flex {

	header, table := NewTransactionsHeader(pages)

	root := tview.NewFlex().SetDirection(tview.FlexRow)
	root.SetBorder(true).SetTitle("View transactions")
	root.AddItem(header, 3, 0, true)
	root.AddItem(table, 0, 1, false)

	return root
}
