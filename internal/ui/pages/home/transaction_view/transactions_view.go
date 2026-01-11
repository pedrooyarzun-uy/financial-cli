package transactionview

import "github.com/rivo/tview"

func NewTransactionsView(app *tview.Application, pages *tview.Pages) *tview.Flex {

	header, transactions := NewTransactionsHeader(pages)
	results := NewTransactionsTable(transactions)

	root := tview.NewFlex().SetDirection(tview.FlexRow)
	root.SetBorder(true).SetTitle("View transactions")
	root.AddItem(header, 1, 0, true)
	root.AddItem(results, 0, 1, false)

	return root
}
