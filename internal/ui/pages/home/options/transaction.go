package options

import (
	"github.com/gdamore/tcell/v2"
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/services"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components/dropdowns"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/validators"
	"github.com/rivo/tview"
)

func NewTransaction(app *tview.Application, pages *tview.Pages) *tview.Flex {
	form := components.NewForm()
	ts := services.NewTransactionService(api.CLIENT)

	//Amount dropdown
	amount := components.NewInputField("Amount:", 10, 17)
	form.AddFormItem(amount)

	//Type dropdown
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
	subcategory := dropdowns.NewSubCategoryDropdown("Subcategory...", 10, 30)

	category := dropdowns.NewCategoryDropdown("Category", 10, 30, func(categoryID int) {
		subcategory.LoadSubCategories(categoryID)
	})

	form.AddFormItem(category)
	form.AddFormItem(subcategory)

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
		_, selectedType := type_.GetCurrentOption()
		_, selectedCurrency := currency.GetCurrentOption()

		categoryID, ok := category.GetSelectedCategoryID()

		if !ok {
			modal := components.NewWarningModal("Please select a valid category", pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		subcategoryID, ok := subcategory.GetSelectedSubCategoryID()

		if !ok {
			modal := components.NewWarningModal("Please select a valid subcategory", pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		notes := form.GetFormItem(5).(*components.TextArea).GetText()

		//Pending, select account and subcategory in form
		err = ts.Add(amount, 10, currencyMap[selectedCurrency], typeMap[selectedType], categoryID, subcategoryID, notes)

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
