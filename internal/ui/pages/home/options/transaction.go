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
	form := components.NewForm()

	cs := services.NewCategoryService(api.CLIENT)
	ts := services.NewTransactionService(api.CLIENT)

	categoryOptions, err := cs.GetAllForDropdown("Add new category...", false)
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
	amount := components.NewInputField("Amount:", 10, 17)
	form.AddFormItem(amount)

	type_ := components.NewDropDown("Type:", 10, 30, []string{"Income", "Outcome"}, nil)
	form.AddFormItem(type_)

	typeMap := map[string]int{
		"Income":  1,
		"Outcome": 2,
	}

	//Currency dropdown
	currency := components.NewDropDown("Currency:", 10, 30, []string{"UY", "USD"}, nil)
	form.AddFormItem(currency)

	currencyMap := map[string]int{
		"USD": 1,
		"UY":  2,
	}

	category := components.NewDropDown("Category: ", 10, 30, categoryLabels, func(text string, index int) {
		if categoryMap[text] == -1 {
			categoryModal := NewCategory(pages)
			pages.AddPage("category", categoryModal, true, true)
		}
	})
	form.AddFormItem(category)

	//text area for notes
	notes := components.NewTextArea("Notes: ", "Add your notes...", 30)
	form.AddFormItem(notes)

	//Back button
	form.AddButton("Go Back", func() {
		pages.SwitchToPage("home")
	})

	//Save button
	form.AddButton("Save", func() {

		//Check amount
		amount, err := validators.CheckAmount(form.GetFormItem(0).(*components.InputField).GetText())

		if err != nil {
			modal := components.NewWarningModal(err.Error(), pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		//Get values from form
		_, type_ := form.GetFormItem(1).(*components.DropDown).GetCurrentOption()
		_, currency := form.GetFormItem(2).(*tview.DropDown).GetCurrentOption()
		_, category := form.GetFormItem(3).(*tview.DropDown).GetCurrentOption()
		categoryID := categoryMap[category]

		notes := form.GetFormItem(4).(*components.TextArea).GetText()

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
	form.SetButtonBackgroundColor(tcell.ColorGray).
		SetButtonTextColor(tcell.ColorWhite)

	//Flex for return
	flex := tview.NewFlex().SetDirection(tview.FlexColumn)
	flex.SetBorder(true).SetTitle("New Transaction").SetTitleAlign(1)
	flex.AddItem(form, 30, 1, true)

	return flex
}
