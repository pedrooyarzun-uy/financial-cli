package transactionview

import (
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/services"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"
	"github.com/rivo/tview"
)

func NewTransactionsHeader(pages *tview.Pages) (*tview.Grid, *tview.Table) {

	currentPage := 1

	ts := services.NewTransactionService(api.CLIENT)
	cs := services.NewCategoryService(api.CLIENT)
	ss := services.NewSubcategoryService(api.CLIENT)

	table := tview.NewTable().SetBorders(true)
	SetHeaders(table)

	transactions, maxPage, _ := ts.GetTransactionsByDetail("", "", 0, 0, currentPage, 10)
	LoadData(table, transactions)

	fromInput := tview.NewInputField().
		SetLabel("From: ").
		SetLabelWidth(6).
		SetFieldWidth(10)

	toInput := tview.NewInputField().
		SetLabel("To: ").
		SetLabelWidth(4).
		SetFieldWidth(10)

	subcategoriesMap := make(map[string]int)

	subcategoryDrop := components.NewDropDown("Subcategory: ", 13, 14, []string{"Select category..."}, nil)

	categories, _ := cs.GetAllForDropdown()
	categoryLabels := make([]string, 0, len(categories))
	categoryMap := make(map[string]int, len(categories))

	for _, opt := range categories {
		categoryLabels = append(categoryLabels, opt.Label)
		categoryMap[opt.Label] = opt.Value
	}

	categoryDrop := components.NewDropDown("Category: ", 10, 14, categoryLabels, func(text string, index int) {
		selctedCategory := categoryMap[text]

		subs, err := ss.GetAllForDropdown(selctedCategory)

		if err != nil {
			subcategoriesMap = make(map[string]int)
			subcategoryDrop.SetOptions([]string{"(error loading)"}, nil)
			subcategoryDrop.SetCurrentOption(0)
			return
		}

		labels := make([]string, 0, len(subs))
		subcategoriesMap = make(map[string]int, len(subs))

		for _, opt := range subs {
			labels = append(labels, opt.Label)
			subcategoriesMap[opt.Label] = opt.Value
		}

		if len(labels) == 0 {
			labels = []string{"(no subcategories)"}
			subcategoriesMap = make(map[string]int)
		}

		subcategoryDrop.SetOptions(labels, nil)
		subcategoryDrop.SetCurrentOption(0)
	})

	backBtn := tview.NewButton("<< Back").
		SetSelectedFunc(func() { pages.SwitchToPage("home") })

	searchBtn := tview.NewButton("Search").SetSelectedFunc(func() {
		categoryDrop.GetCurrentOption()
		_, categoryText := categoryDrop.GetCurrentOption()
		_, subcategoryText := subcategoryDrop.GetCurrentOption()

		categoryID := categoryMap[categoryText]
		subcategoryID := subcategoriesMap[subcategoryText]

		from := fromInput.GetText()
		to := toInput.GetText()

		transactions, _, _ := ts.GetTransactionsByDetail(from, to, categoryID, subcategoryID, currentPage, 10)
		RefreshTable(table)
		SetHeaders(table)
		LoadData(table, transactions)
	})

	nextBtn := tview.NewButton("Next page >>").
		SetSelectedFunc(func() {
			if currentPage < maxPage {
				currentPage += 1
			}

			categoryDrop.GetCurrentOption()
			_, categoryText := categoryDrop.GetCurrentOption()
			_, subcategoryText := subcategoryDrop.GetCurrentOption()

			categoryID := categoryMap[categoryText]
			subcategoryID := subcategoriesMap[subcategoryText]

			from := fromInput.GetText()
			to := toInput.GetText()

			transactions, _, _ := ts.GetTransactionsByDetail(from, to, categoryID, subcategoryID, currentPage, 10)
			RefreshTable(table)
			SetHeaders(table)
			LoadData(table, transactions)
		})

	prevBtn := tview.NewButton("<< Previous page").
		SetSelectedFunc(func() {
			if currentPage > 1 {
				currentPage -= 1
			}
			categoryDrop.GetCurrentOption()
			_, categoryText := categoryDrop.GetCurrentOption()
			_, subcategoryText := subcategoryDrop.GetCurrentOption()

			categoryID := categoryMap[categoryText]
			subcategoryID := subcategoriesMap[subcategoryText]

			from := fromInput.GetText()
			to := toInput.GetText()

			transactions, _, _ := ts.GetTransactionsByDetail(from, to, categoryID, subcategoryID, currentPage, 10)
			RefreshTable(table)
			SetHeaders(table)
			LoadData(table, transactions)
		})

	top := tview.NewGrid().
		SetRows(1, 1).
		SetColumns(12, -2, -2, -3, -3, 12)
	top.SetGap(1, 1)

	// row 0: filters
	top.AddItem(backBtn, 0, 0, 1, 1, 0, 0, false)
	top.AddItem(fromInput, 0, 1, 1, 1, 0, 0, true)
	top.AddItem(toInput, 0, 2, 1, 1, 0, 0, false)
	top.AddItem(categoryDrop, 0, 3, 1, 1, 0, 0, false)
	top.AddItem(subcategoryDrop, 0, 4, 1, 1, 0, 0, false)
	top.AddItem(searchBtn, 0, 5, 1, 1, 0, 0, false)

	//row 1: paging
	top.AddItem(prevBtn, 1, 0, 1, 3, 0, 0, false)
	top.AddItem(nextBtn, 1, 3, 1, 3, 0, 0, false)

	return top, table
}
