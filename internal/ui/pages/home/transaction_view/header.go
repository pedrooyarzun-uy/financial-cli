package transactionview

import "github.com/rivo/tview"

func NewTransactionsHeader(pages *tview.Pages) *tview.Grid {
	fromInput := tview.NewInputField().
		SetLabel("From: ").
		SetLabelWidth(6).
		SetFieldWidth(10)

	toInput := tview.NewInputField().
		SetLabel("To: ").
		SetLabelWidth(4).
		SetFieldWidth(10)

	categoryDrop := tview.NewDropDown().
		SetLabel("Category: ").
		SetLabelWidth(10).
		SetFieldWidth(14).
		SetOptions([]string{"Prueba"}, nil)

	subcategoryDrop := tview.NewDropDown().
		SetLabel("Subcategory: ").
		SetLabelWidth(13).
		SetFieldWidth(14).
		SetOptions([]string{"Prueba"}, nil)

	backBtn := tview.NewButton("⬅ Back").
		SetSelectedFunc(func() { pages.SwitchToPage("home") })

	searchBtn := tview.NewButton("Search 🔎")

	top := tview.NewGrid().
		SetRows(2).
		SetColumns(10, 18, 16, 26, 30, 12)
	top.SetGap(1, 1)

	top.AddItem(backBtn, 0, 0, 1, 1, 0, 0, false)
	top.AddItem(fromInput, 0, 1, 1, 1, 0, 0, true)
	top.AddItem(toInput, 0, 2, 1, 1, 0, 0, false)
	top.AddItem(categoryDrop, 0, 3, 1, 1, 0, 0, false)
	top.AddItem(subcategoryDrop, 0, 4, 1, 1, 0, 0, false)
	top.AddItem(searchBtn, 0, 5, 1, 1, 0, 0, false)

	return top
}
