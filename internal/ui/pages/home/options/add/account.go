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
	name := components.NewInputField("Name: [red::b]*", 10, 30)
	form.AddFormItem(name)

	//Number Input
	number := components.NewInputField("Number: ", 10, 30)
	form.AddFormItem(number)

	//Bank Dropdown
	bank := dropdowns.NewBankDropdown("Bank: [red::b]*", 10, 30)
	form.AddFormItem(bank)

	//Currency Dropdown
	currency := dropdowns.NewCurrencyDropdown("Currency: [red::b]*", 10, 30)
	form.AddFormItem(currency)

	form.AddButton("Go Back", func() {
		pages.SwitchToPage("add_page")
	})
	form.AddButton("Add Account", nil).SetButtonBackgroundColor(tcell.ColorGreenYellow)

	form.SetButtonTextColor(tcell.ColorWhite).SetButtonBackgroundColor(tcell.ColorGray)

	titleText := tview.NewTextView().
		SetText("[yellow::b]Enter account data\n[white::b]Fields with [red::b]* [white::b]are mandatories").
		SetTextAlign(tview.AlignCenter)
	titleText.SetDynamicColors(true)

	centeredForm := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(form, 60, 0, true).
		AddItem(nil, 0, 1, false)

	content := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(titleText, 2, 0, false).
		AddItem(nil, 1, 0, false).
		AddItem(centeredForm, 0, 1, true)

	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.SetBorder(true).SetTitle("Add Account").SetTitleAlign(tview.AlignCenter)
	flex.AddItem(nil, 0, 1, false)
	flex.AddItem(content, 14, 0, true)
	flex.AddItem(nil, 0, 1, false)

	return flex
}
