package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	p "github.com/pedrooyarzun-uy/financial-cli/internal/ui/pages"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/pages/home"
	"github.com/rivo/tview"
)

func main() {

	api.Init("http://localhost:8080")

	app := tview.NewApplication()
	pages := tview.NewPages()

	login := p.NewLoginPage(app, pages)
	signUp := p.NewSignUp(pages)
	home := home.NewHomePage(app, pages)

	pages.AddPage("login", login, true, true)
	pages.AddPage("home", home, true, false)
	pages.AddPage("sign_up", signUp, true, false)

	app.SetBeforeDrawFunc(func(screen tcell.Screen) bool {
		screen.Clear()
		screen.Fill(' ', tcell.StyleDefault.Background(tcell.ColorBlack))
		return false
	})

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
