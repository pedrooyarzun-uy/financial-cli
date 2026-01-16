package transactionview

import (
	"strconv"
	"strings"

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
		//Parsing color for tcell
		val.Color = strings.TrimPrefix(val.Color, "#")
		v, _ := strconv.ParseInt(val.Color, 16, 32)

		idx += 1
		table.SetCell(idx, 0, tview.NewTableCell(val.Category).SetTextColor(tcell.NewHexColor(int32(v))))
		table.SetCell(idx, 1, tview.NewTableCell(val.Subcategory))

		if val.Type == 1 {
			table.SetCell(idx, 2, tview.NewTableCell(strconv.FormatFloat(val.Amount, 'f', 2, 64)).SetTextColor(tcell.ColorGreen))
		} else if val.Type == 2 {
			table.SetCell(idx, 2, tview.NewTableCell(strconv.FormatFloat(val.Amount, 'f', 2, 64)).SetTextColor(tcell.ColorRed))
		}

		table.SetCell(idx, 3, tview.NewTableCell(val.Currency))
		table.SetCell(idx, 4, tview.NewTableCell(val.Notes))
		table.SetCell(idx, 5, tview.NewTableCell(val.Date.String()))
	}

	return table
}
