package add

import (
	"fmt"

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

	//Account dropdown
	account := dropdowns.NewAccountDropdown("Account:", 10, 30)
	form.AddFormItem(account)

	//Type dropdown
	type_ := dropdowns.NewTypeDropdown("Type:", 10, 30)
	form.AddFormItem(type_)

	//Currency dropdown
	currency := dropdowns.NewCurrencyDropdown("Currency:", 10, 30)
	form.AddFormItem(currency)

	//Subcategory dropdown
	subcategory := dropdowns.NewSubCategoryDropdown("Subcategory:", 10, 30)

	//Category dropdown
	category := dropdowns.NewCategoryDropdown("Category:", 10, 30, func(categoryID int) {
		subcategory.LoadSubCategories(categoryID)
	})
	form.AddFormItem(category)
	form.AddFormItem(subcategory)

	//Notes textarea
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

		//TypeID
		typeID, ok := type_.GetSelectedTypeID()
		if !ok {
			modal := components.NewWarningModal("Please select a valid type of transaction", pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		//CategoryID
		categoryID, ok := category.GetSelectedCategoryID()
		if !ok {
			fmt.Println(categoryID, ok)
			modal := components.NewWarningModal("Please select a valid category", pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		//SubcategoryID
		subcategoryID, ok := subcategory.GetSelectedSubCategoryID()
		if !ok {
			modal := components.NewWarningModal("Please select a valid subcategory", pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		//AccountID
		accountID, ok := account.GetSelectedAccountID()
		if !ok {
			modal := components.NewWarningModal("Please select a valid account", pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		//CurrencyID
		currencyID, ok := currency.GetSelectedCurrencyID()
		if !ok {
			modal := components.NewWarningModal("Please select a valid currency", pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		//Notes value
		notes := form.GetFormItem(6).(*components.TextArea).GetText()

		//Send to backend
		err = ts.Add(amount, accountID, currencyID, typeID, categoryID, subcategoryID, notes)

		if err != nil {
			modal := components.NewWarningModal(err.Error(), pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		modal := components.NewSuccessModal("Transaction saved correctly!", pages)
		pages.AddPage("modal", modal, true, true)
	})

	//Render form
	form.SetButtonBackgroundColor(tcell.ColorGray).
		SetButtonTextColor(tcell.ColorWhite)

	flex := tview.NewFlex().SetDirection(tview.FlexColumn)
	flex.SetBorder(true).SetTitle("New Transaction").SetTitleAlign(1)
	flex.AddItem(form, 30, 1, true)

	return flex
}
