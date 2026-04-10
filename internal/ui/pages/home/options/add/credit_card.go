package add

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/helpers"
	"github.com/pedrooyarzun-uy/financial-cli/internal/services"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components/dropdowns"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/validators"
	"github.com/rivo/tview"
)

func NewAddCreditCard(app *tview.Application, pages *tview.Pages) *tview.Flex {
	ccs := services.NewCreditCardService(api.CLIENT)

	form := components.NewForm()

	//Name input
	name := components.NewInputField("Name: [red::b]*", 10, 30)
	form.AddFormItem(name)

	//Bank dropdown
	bank := dropdowns.NewBankDropdown("Bank: [red::b]*", 10, 30)
	form.AddFormItem(bank)

	//Close Day input
	closeDay := components.NewInputField("Close Day: [red::b]*", 10, 30)
	form.AddFormItem(closeDay)

	//Due Day input
	dueDay := components.NewInputField("Due Day: [red::b]*", 10, 30)
	form.AddFormItem(dueDay)

	//Credit Limit input
	creditLimit := components.NewInputField("Credit Limit: [red::b]*", 10, 30)
	form.AddFormItem(creditLimit)

	form.AddButton("Go Back", func() {
		pages.SwitchToPage("add_page")
	})

	form.AddButton("Add Credit Card", func() {
		name := form.GetFormItem(0).(*components.InputField).GetText()
		bankID, ok := bank.GetSelectedBankID()

		if !ok {
			modal := components.NewWarningModal("Please select a valid bank", pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		closeDay, ok := helpers.StringToInt(form.GetFormItem(2).(*components.InputField).GetText())

		if !ok || closeDay < 1 || closeDay > 31 {
			modal := components.NewWarningModal("Close Day must be a number between 1 and 31", pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		dueDay, ok := helpers.StringToInt(form.GetFormItem(3).(*components.InputField).GetText())

		if !ok || dueDay < 1 || dueDay > 31 {
			modal := components.NewWarningModal("Due Day must be a number between 1 and 31", pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		creditLimit, err := validators.CheckAmount(form.GetFormItem(4).(*components.InputField).GetText())

		if err != nil {
			modal := components.NewWarningModal("Limit must be a number (float)", pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		err = ccs.Add(name, bankID, closeDay, dueDay, creditLimit)

		if err != nil {
			modal := components.NewWarningModal(err.Error(), pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		modal := components.NewSuccessModal("Credit Card saved correctly!", pages)
		pages.AddPage("modal", modal, true, true)
	})

	form.SetButtonTextColor(tcell.ColorWhite).SetButtonBackgroundColor(tcell.ColorGray)

	titleText := tview.NewTextView().
		SetText("[yellow::b]Enter credit card data\n[white]Fields with [red::b]* [white]are mandatories").
		SetTextAlign(tview.AlignCenter)
	titleText.SetDynamicColors(true)

	preview := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true)
	preview.SetBorder(true).SetTitle("Credit Card Preview")

	tips := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetText(
			"[yellow::b]Tips\n\n" +
				"[white]- Name and Bank are required\n" +
				"[white]- Close and Due Day are required\n" +
				"[white]- Credit Limit is mandatory\n" +
				"[white]- Use a clear name, e.g.:\n" +
				"  [green]Savings Account[-], [green]Payroll[-], [green]Cash[-]\n",
		)
	tips.SetBorder(true).SetTitle("Help")

	preview.SetText(fmt.Sprintf(
		"[yellow::b]Credit Card Preview\n\n" +
			"[white]Name: [green]\n" +
			"[white]Bank: [green]\n" +
			"[white]Close Day: [green]\n" +
			"[white]Due Day: [green]\n" +
			"[white]Credit Limit: [green]\n",
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
	flex.SetBorder(true).SetTitle("Add Credit Card").SetTitleAlign(tview.AlignCenter)
	flex.AddItem(nil, 0, 1, false)
	flex.AddItem(content, 22, 0, true)
	flex.AddItem(nil, 0, 1, false)

	return flex
}
