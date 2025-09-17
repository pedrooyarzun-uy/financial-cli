package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"
	"github.com/rivo/tview"
)

func LoginPage(app *tview.Application, pages *tview.Pages) *tview.Flex {
	email := components.NewInput("Email: ", 30, tcell.ColorGreen, tcell.ColorBlack, false)
	password := components.NewInput("Password: ", 30, tcell.ColorGreen, tcell.ColorBlack, true)

	email.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			app.SetFocus(password)
		}
	})

	password.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			if password.GetText() == "admin" {
				pages.SwitchToPage("home")
			}
		}
	})

	box := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText("Welcome to Financial CLI!").SetTextColor(tcell.ColorGreen), 2, 0, false).
		AddItem(email, 1, 0, true).
		AddItem(password, 1, 0, false)

	return box
}
