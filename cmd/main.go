package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui"
	"github.com/rivo/tview"
)

func main() {
	login := ui.LoginPage()

	app := tview.NewApplication()

	app.SetBeforeDrawFunc(func(screen tcell.Screen) bool {
		screen.Clear()
		screen.Fill(' ', tcell.StyleDefault.Background(tcell.ColorBlack))
		return false
	})

	if err := app.SetRoot(login, true).Run(); err != nil {
		panic(err)
	}
}
