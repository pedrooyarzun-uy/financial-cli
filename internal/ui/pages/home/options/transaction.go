package options

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewTransaction(app *tview.Application, pages *tview.Pages) *tview.Flex {
	form := tview.NewForm()

	categories := []string{
		"Food", "Clothes",
		"Transport", "Education",
		"Health", "Create New one",
	}

	form.AddInputField("Amount:", "$", 30, nil, nil).SetLabelColor(tcell.ColorCornflowerBlue).SetFieldBackgroundColor(tcell.ColorDarkSlateGray)
	form.AddDropDown("Currency:", []string{"UY", "USD"}, 0, nil)
	form.AddDropDown("Type:", []string{"Income", "Expense"}, 0, nil)
	form.AddDropDown("Category:", categories, 1, func(option string, optionIndex int) {
		if optionIndex == 5 {
			panic("Hola mundo")
		}
	})
	form.AddTextArea("Notes:", "Add your notes...", 30, 4, 30, nil)

	flex := tview.NewFlex().SetDirection(tview.FlexColumn)
	flex.SetBorder(true).SetTitle("NEW TRANSACTION").SetBorderColor(tcell.ColorDarkGreen).SetTitleAlign(0).SetTitleColor(tcell.ColorDarkOliveGreen)

	flex.AddItem(form, 30, 1, true)

	return flex
}
