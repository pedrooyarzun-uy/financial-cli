package transactionview

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/pedrooyarzun-uy/financial-cli/internal/dto"
	"github.com/rivo/tview"
)

func NewTransactionsTable(transactions []dto.TransactionByDetail) *tview.Table {

	table := tview.NewTable().SetBorders(true)

	table.SetCell(0, 0, tview.NewTableCell("Category").SetSelectable(false).SetAttributes(tcell.AttrBold).SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 1, tview.NewTableCell("Subcategory").SetSelectable(false).SetAttributes(tcell.AttrBold).SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 2, tview.NewTableCell("Amount").SetSelectable(false).SetAttributes(tcell.AttrBold).SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 3, tview.NewTableCell("Currency").SetSelectable(false).SetAttributes(tcell.AttrBold).SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 4, tview.NewTableCell("Notes").SetSelectable(false).SetAttributes(tcell.AttrBold).SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 5, tview.NewTableCell("Date").SetSelectable(false).SetAttributes(tcell.AttrBold).SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))

	for idx, val := range transactions {
		idx += 1
		table.SetCell(idx, 0, tview.NewTableCell(val.Category))
		table.SetCell(idx, 1, tview.NewTableCell(val.Subcategory))
		table.SetCell(idx, 2, tview.NewTableCell(strconv.FormatFloat(val.Amount, 'f', 2, 64)))
		table.SetCell(idx, 3, tview.NewTableCell(val.Currency))
		table.SetCell(idx, 4, tview.NewTableCell(val.Notes))
		table.SetCell(idx, 5, tview.NewTableCell(val.Date.String()))
	}

	return table
}
