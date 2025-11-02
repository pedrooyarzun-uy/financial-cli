package pages

import (
	"github.com/gdamore/tcell/v2"
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/models"
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
		email := form.GetFormItem(0).(*tview.InputField).GetText()
		password := form.GetFormItem(1).(*tview.InputField).GetText()

		req := models.SignInReq{
			Email:    email,
			Password: password,
		}

		res := models.SignInRes{}

		err := api.CLIENT.PostMethod("/user/sign-in", &res, req, true)

		if err != nil {
			form.AddTextView("Error", err.Error(), 30, 1, false, false)
		} else {
			form.AddTextView("Response", res.Auth+" "+res.Message, 30, 1, false, false)
		}

	})

	form.AddButton("Sign up", func() {
		pages.SwitchToPage("sign_up")
	})

	form.AddTextView("Don't have an account yet?", "", 1, 1, false, false)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true)

	return flex
}
