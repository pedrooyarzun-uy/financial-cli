package main

import (
	"github.com/gdamore/tcell/v2"
	p "github.com/pedrooyarzun-uy/financial-cli/internal/ui/pages"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/pages/home"
	"github.com/rivo/tview"
)

func main() {

	app := tview.NewApplication()
	pages := tview.NewPages()

	login := p.NewLoginPage(app, pages)
	home := home.NewHomePage(app, pages)

	pages.AddPage("login", login, true, true)
	pages.AddPage("home", home, true, false)

	app.SetBeforeDrawFunc(func(screen tcell.Screen) bool {
		screen.Clear()
		screen.Fill(' ', tcell.StyleDefault.Background(tcell.ColorBlack))
		return false
	})

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
