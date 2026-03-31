package add

import (
	"github.com/gdamore/tcell/v2"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components/dropdowns"
	"github.com/rivo/tview"
)

func NewAddAccount(app *tview.Application, pages *tview.Pages) *tview.Flex {
	form := components.NewForm()
	//as := services.NewAccountService(api.CLIENT)

	//Name Input
	name := components.NewInputField("Name:", 10, 17)
	form.AddFormItem(name)

	//Number Input
	number := components.NewInputField("Number:", 10, 17)
	form.AddFormItem(number)

	//Bank Dropdown
	bank := dropdowns.NewBankDropdown("Bank", 10, 30)
	form.AddFormItem(bank)

	//Currency Dropdown
	currency := dropdowns.NewCurrencyDropdown("Currency:", 10, 30)
	form.AddFormItem(currency)

	form.SetButtonBackgroundColor(tcell.ColorGray).
		SetButtonTextColor(tcell.ColorWhite)

	flex := tview.NewFlex().SetDirection(tview.FlexColumn)
	flex.SetBorder(true).SetTitle("Add Account").SetTitleAlign(1)
	flex.AddItem(form, 30, 1, true)

	return flex
}
