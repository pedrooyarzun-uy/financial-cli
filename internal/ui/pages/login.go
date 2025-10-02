package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewLoginPage(app *tview.Application, pages *tview.Pages) *tview.Flex {

	form := tview.NewForm()

	//Styles for form
	form.SetFieldBackgroundColor(tcell.ColorPaleVioletRed)
	form.SetLabelColor(tcell.ColorWhite)
	form.SetButtonBackgroundColor(tcell.ColorPaleVioletRed)

	form.AddInputField("Email:", "", 30, nil, nil).SetFieldBackgroundColor(tcell.ColorPaleVioletRed)
	form.AddPasswordField("Password", "", 30, '*', nil)

	form.AddButton("Sign In", func() {
		text := form.GetFormItem(1).(*tview.InputField).GetText()

		if text == "admin" {
			pages.SwitchToPage("home")
		}
	})

	form.AddButton("Sign up", nil)

	form.AddTextView("Don't have an account yet?", "", 1, 1, false, false)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true)

	return flex
}
