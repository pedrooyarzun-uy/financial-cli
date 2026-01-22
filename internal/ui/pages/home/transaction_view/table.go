package transactionview

import (
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/pedrooyarzun-uy/financial-cli/internal/dto"
	"github.com/rivo/tview"
)

func SetHeaders(table *tview.Table) {

	table.SetCell(0, 0, tview.NewTableCell("Category").SetSelectable(false).SetAttributes(tcell.AttrBold).SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 1, tview.NewTableCell("Subcategory").SetSelectable(false).SetAttributes(tcell.AttrBold).SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 2, tview.NewTableCell("Currency").SetSelectable(false).SetAttributes(tcell.AttrBold).SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 3, tview.NewTableCell("Amount").SetSelectable(false).SetAttributes(tcell.AttrBold).SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 4, tview.NewTableCell("Date").SetSelectable(false).SetAttributes(tcell.AttrBold).SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 5, tview.NewTableCell("Notes").SetSelectable(false).SetAttributes(tcell.AttrBold).SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
}

func RefreshTable(table *tview.Table) {
	//Clean table
	table.Clear()
}

func LoadData(table *tview.Table, transactions []dto.TransactionByDetail) {
	for idx, val := range transactions {
		//Parsing color for tcell
		val.Color = strings.TrimPrefix(val.Color, "#")
		v, _ := strconv.ParseInt(val.Color, 16, 32)

		idx += 1
		table.SetCell(idx, 0, tview.NewTableCell(val.Category).SetTextColor(tcell.NewHexColor(int32(v))))
		table.SetCell(idx, 1, tview.NewTableCell(val.Subcategory))
		table.SetCell(idx, 2, tview.NewTableCell(val.Currency))

		if val.Type == 1 {
			table.SetCell(idx, 3, tview.NewTableCell(strconv.FormatFloat(val.Amount, 'f', 2, 64)).SetTextColor(tcell.ColorGreen))
		} else if val.Type == 2 {
			table.SetCell(idx, 3, tview.NewTableCell(strconv.FormatFloat(val.Amount, 'f', 2, 64)).SetTextColor(tcell.ColorRed))
		}

		table.SetCell(idx, 4, tview.NewTableCell(val.Date.Format("2006-01-02 15:04:05")))
		table.SetCell(idx, 5, tview.NewTableCell(val.Notes).SetExpansion(3))
	}
}
