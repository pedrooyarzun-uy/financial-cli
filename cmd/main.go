package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui"
	"github.com/rivo/tview"
)

func main() {

	app := tview.NewApplication()
	pages := tview.NewPages()

	login := ui.LoginPage(app, pages)
	home := ui.HomePage(app, pages)

	pages.AddPage("login", login, true, true)
	pages.AddPage("home", home, true, false)

	app.SetBeforeDrawFunc(func(screen tcell.Screen) bool {
		screen.Clear()
		screen.Fill(' ', tcell.StyleDefault.Background(tcell.ColorBlack))
		return false
	})

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
