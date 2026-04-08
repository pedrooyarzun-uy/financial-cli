package add

import (
	"github.com/gdamore/tcell/v2"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components/dropdowns"
	"github.com/rivo/tview"
)

func NewAddCreditCard(app *tview.Application, pages *tview.Pages) *tview.Flex {
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

	form.AddButton("Add Credit Card", nil)

	form.SetButtonTextColor(tcell.ColorWhite).SetButtonBackgroundColor(tcell.ColorGray)

	leftPanel := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true)

	return leftPanel
}
