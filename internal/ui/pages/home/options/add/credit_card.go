package add

import (
	"fmt"

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
