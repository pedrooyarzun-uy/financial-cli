package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"
	"github.com/rivo/tview"
)

func NewLoginPage(app *tview.Application, pages *tview.Pages) *tview.Flex {
	email := components.NewInput("Email: ", 30, tcell.ColorGreen, tcell.ColorBlack, false)
	password := components.NewInput("Password: ", 30, tcell.ColorGreen, tcell.ColorBlack, true)

	form := tview.NewForm()

	form.AddFormItem(email)
	form.AddFormItem(password)

	form.AddButton("Sign In", func() {
		if password.GetText() == "admin" {
			pages.SwitchToPage("home")
		}
	})

	form.AddTextView("Don't have an account yet?", "", 1, 1, false, false)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true)

	return flex
}
