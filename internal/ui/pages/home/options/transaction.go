package options

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/validators"
	"github.com/rivo/tview"
)

func NewTransaction(app *tview.Application, pages *tview.Pages) *tview.Flex {
	form := tview.NewForm()

	categories := []string{
		"Food", "Clothes",
		"Transport", "Education",
		"Health", "Create New one",
	}

	//Form
	form.AddInputField("Amount:", "", 30, nil, nil).
		SetLabelColor(tcell.ColorCornflowerBlue).
		SetFieldBackgroundColor(tcell.ColorDarkSlateGray)
	form.AddDropDown("Type:", []string{"Income", "Expense"}, 0, nil)
	form.AddDropDown("Currency:", []string{"UY", "USD"}, 0, nil)
	form.AddDropDown("Category:", categories, 1, nil)
	form.AddTextArea("Notes:", "Add your notes...", 30, 4, 30, nil)

	//Back button
	form.AddButton("Go Back", nil).
		SetButtonBackgroundColor(tcell.ColorLightGoldenrodYellow).
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
		notes := form.GetFormItem(4).(*tview.TextArea).GetText()

		fmt.Println(type_, currency, category, notes, amount)

	})
	saveBtn.SetButtonBackgroundColor(tcell.ColorDarkGreen).
		SetButtonTextColor(tcell.ColorWhite)

	//Flex for return
	flex := tview.NewFlex().SetDirection(tview.FlexColumn)
	flex.SetBorder(true).SetTitle("NEW TRANSACTION").SetBorderColor(tcell.ColorDarkGreen).SetTitleAlign(0).SetTitleColor(tcell.ColorDarkOliveGreen)
	flex.AddItem(form, 30, 1, true)

	return flex
}
