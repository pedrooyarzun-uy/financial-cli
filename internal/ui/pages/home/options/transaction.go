package options

import (
	"github.com/gdamore/tcell/v2"
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/services"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/validators"
	"github.com/rivo/tview"
)

func NewTransaction(app *tview.Application, pages *tview.Pages) *tview.Flex {
	form := tview.NewForm()

	cs := services.NewCategoryService(api.CLIENT)
	ts := services.NewTransactionService(api.CLIENT)

	categoryOptions, err := cs.GetAllForDropdown()
	if err != nil {
		form.AddTextView("Error", err.Error(), 30, 1, false, false)
	}

	categoryLabels := make([]string, 0, len(categoryOptions))
	categoryMap := make(map[string]int, len(categoryOptions))

	for _, opt := range categoryOptions {
		categoryLabels = append(categoryLabels, opt.Label)
		categoryMap[opt.Label] = opt.Value
	}

	//TODO: add error modal with timeout and redirect to home
	if err != nil {
		form.AddTextView("Error", err.Error(), 30, 1, false, false)
	}

	//Form
	form.AddInputField("Amount:", "", 30, nil, nil).
		SetLabelColor(tcell.ColorCornflowerBlue).
		SetFieldBackgroundColor(tcell.ColorDarkSlateGray)
	form.AddDropDown("Type:", []string{"Income", "Expense"}, 0, nil)

	typeMap := map[string]int{
		"Income":  1,
		"Expense": 2,
	}

	form.AddDropDown("Currency:", []string{"UY", "USD"}, 0, nil)

	currencyMap := map[string]int{
		"USD": 1,
		"UY":  2,
	}

	form.AddDropDown("Category:", categoryLabels, 0, func(option string, optionIndex int) {
		if categoryMap[option] == -1 {
			categoryModal := NewCategory(pages)
			pages.AddPage("category", categoryModal, true, true)
		}
	})

	form.AddTextArea("Notes:", "Add your notes...", 30, 4, 30, nil)

	//Back button
	backBtn := form.AddButton("Go Back", func() {
		pages.SwitchToPage("home")
	})

	backBtn.SetButtonBackgroundColor(tcell.ColorLightGoldenrodYellow).
		SetButtonTextColor(tcell.ColorBlack)

	//Save form and validate
	saveBtn := form.AddButton("Save", func() {

		//Check amount
		amount, err := validators.CheckAmount(form.GetFormItem(0).(*tview.InputField).GetText())

		if err != nil {
			modal := components.NewWarningModal(err.Error(), pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		//Get values from form
		_, type_ := form.GetFormItem(1).(*tview.DropDown).GetCurrentOption()
		_, currency := form.GetFormItem(2).(*tview.DropDown).GetCurrentOption()
		_, category := form.GetFormItem(3).(*tview.DropDown).GetCurrentOption()
		categoryID := categoryMap[category]

		notes := form.GetFormItem(4).(*tview.TextArea).GetText()

		//Pending, select account and subcategory in form
		err = ts.Add(amount, 2, currencyMap[currency], typeMap[type_], categoryID, notes)

		if err != nil {
			modal := components.NewWarningModal(err.Error(), pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		modal := components.NewSuccessModal("Transaction saved correctly!", pages)
		pages.AddPage("modal", modal, true, true)

	})
	saveBtn.SetButtonBackgroundColor(tcell.ColorDarkGreen).
		SetButtonTextColor(tcell.ColorWhite)

	//Flex for return
	flex := tview.NewFlex().SetDirection(tview.FlexColumn)
	flex.SetBorder(true).SetTitle("NEW TRANSACTION").SetBorderColor(tcell.ColorDarkGreen).SetTitleAlign(0).SetTitleColor(tcell.ColorDarkOliveGreen)
	flex.AddItem(form, 30, 1, true)

	return flex
}
