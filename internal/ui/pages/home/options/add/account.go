package add

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/services"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components/dropdowns"
	"github.com/rivo/tview"
)

func NewAddAccount(app *tview.Application, pages *tview.Pages) *tview.Flex {
	form := components.NewForm()
	as := services.NewAccountService(api.CLIENT)

	// Name Input
	name := components.NewInputField("Name: [red::b]*", 10, 30)
	form.AddFormItem(name)

	// Number Input
	number := components.NewInputField("Number: ", 10, 30)
	form.AddFormItem(number)

	// Bank Dropdown
	bank := dropdowns.NewBankDropdown("Bank: [red::b]*", 10, 30)
	form.AddFormItem(bank)

	// Currency Dropdown
	currency := dropdowns.NewCurrencyDropdown("Currency: [red::b]*", 10, 30)
	form.AddFormItem(currency)

	form.AddButton("Go Back", func() {
		pages.SwitchToPage("add_page")
	})

	form.AddButton("Add Account", func() {
		name := form.GetFormItem(0).(*components.InputField).GetText()
		number := form.GetFormItem(1).(*components.InputField).GetText()

		bankID, ok := bank.GetSelectedBankID()
		if !ok {
			modal := components.NewWarningModal("Please select a valid bank", pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		currencyID, ok := currency.GetSelectedCurrencyID()
		if !ok {
			modal := components.NewWarningModal("Please select a valid currency", pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		err := as.Add(name, number, currencyID, bankID)

		if err != nil {
			modal := components.NewWarningModal(err.Error(), pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		modal := components.NewSuccessModal("Account added succesfully!", pages)
		pages.AddPage("modal", modal, true, true)

	})

	form.SetButtonTextColor(tcell.ColorWhite).SetButtonBackgroundColor(tcell.ColorGray)

	titleText := tview.NewTextView().
		SetText("[yellow::b]Enter account data\n[white]Fields with [red::b]* [white]are mandatories").
		SetTextAlign(tview.AlignCenter)
	titleText.SetDynamicColors(true)

	preview := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true)
	preview.SetBorder(true).SetTitle("Account Preview")

	tips := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetText(
			"[yellow::b]Tips\n\n" +
				"[white]- Name and Bank are required\n" +
				"[white]- Currency is required\n" +
				"[white]- Number is optional\n" +
				"[white]- Use a clear name, e.g.:\n" +
				"  [green]Savings Account[-], [green]Payroll[-], [green]Cash[-]\n",
		)
	tips.SetBorder(true).SetTitle("Help")

	preview.SetText(fmt.Sprintf(
		"[yellow::b]Account Preview\n\n" +
			"[white]Name: [green]\n" +
			"[white]Number: [green]\n" +
			"[white]Bank: [green]\n" +
			"[white]Currency: [green]\n\n" +
			"[white]Status: ",
	))

	leftPanel := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true)

	rightPanel := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(preview, 0, 3, false).
		AddItem(nil, 1, 0, false).
		AddItem(tips, 0, 2, false)

	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(leftPanel, 0, 2, true).
		AddItem(nil, 2, 0, false).
		AddItem(rightPanel, 0, 1, false)

	content := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(titleText, 2, 0, false).
		AddItem(nil, 1, 0, false).
		AddItem(body, 0, 1, true)

	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.SetBorder(true).SetTitle("Add Account").SetTitleAlign(tview.AlignCenter)
	flex.AddItem(nil, 0, 1, false)
	flex.AddItem(content, 22, 0, true)
	flex.AddItem(nil, 0, 1, false)

	return flex
}
