package transactionview

import "github.com/rivo/tview"

func NewTransactionsView(app *tview.Application, pages *tview.Pages) *tview.Flex {

	results := NewTransactionsTable()
	top := NewTransactionsHeader(pages)

	root := tview.NewFlex().SetDirection(tview.FlexRow)
	root.SetBorder(true).SetTitle("View transactions")
	root.AddItem(top, 1, 0, true)
	root.AddItem(results, 0, 1, false)

	return root
}
