package transactionview

import (
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/dto"
	"github.com/pedrooyarzun-uy/financial-cli/internal/services"
	"github.com/rivo/tview"
)

func NewTransactionsHeader(pages *tview.Pages) (*tview.Grid, []dto.TransactionByDetail) {

	ts := services.NewTransactionService(api.CLIENT)
	cs := services.NewCategoryService(api.CLIENT)

	categories, _ := cs.GetAllForDropdown()
	categoryLabels := make([]string, 0, len(categories))
	categoryMap := make(map[string]int, len(categories))

	for _, opt := range categories {
		categoryLabels = append(categoryLabels, opt.Label)
		categoryMap[opt.Label] = opt.Value
	}

	transactions, _ := ts.GetTransactionsByDetail("", "", 0, 0)

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
		SetOptions(categoryLabels, nil)

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

	return top, transactions
}
