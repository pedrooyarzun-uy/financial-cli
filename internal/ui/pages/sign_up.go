package pages

import (
	"github.com/gdamore/tcell/v2"
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/models"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/validators"
	"github.com/rivo/tview"
)

func NewSignUp(pages *tview.Pages) *tview.Flex {

	form := tview.NewForm()

	//Styles for form
	form.SetFieldBackgroundColor(tcell.ColorPaleVioletRed)
	form.SetLabelColor(tcell.ColorWhite)
	form.SetButtonBackgroundColor(tcell.ColorPaleVioletRed)

	//Content
	form.AddInputField("Name:", "", 30, nil, nil)
	form.AddInputField("Email:", "", 30, nil, nil)
	form.AddPasswordField("Password:", "", 30, '*', nil)
	form.AddPasswordField("Repeat Password:", "", 30, '*', nil)

	form.AddButton("Sign Up", func() {

		//Get values
		name := form.GetFormItem(0).(*tview.InputField).GetText()
		email := form.GetFormItem(1).(*tview.InputField).GetText()
		password := form.GetFormItem(2).(*tview.InputField).GetText()
		repeatedPasssword := form.GetFormItem(3).(*tview.InputField).GetText()

		if password != repeatedPasssword {
			modal := components.NewWarningModal("Passwords do not match", pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		if !validators.CheckEmail(email) {
			modal := components.NewWarningModal("Email is not valid.", pages)
			pages.AddPage("modal", modal, true, true)
			return
		}

		req := models.SignUpReq{
			Name:     name,
			Email:    email,
			Password: password,
		}

		res := models.SignUpRes{}

		err := api.CLIENT.PostMethod("/user/sign-up", &res, req, true)

		if err != nil {
			//Upgrade: add modal for errors
			form.AddTextView("Error", err.Error(), 30, 1, false, false)
		}

		//Upgrade: add modal and then redirect to login
		pages.SwitchToPage("login")
	})

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexColumn)
	flex.SetBorder(true).
		SetTitle("SIGN UP").
		SetBorderColor(tcell.ColorPaleVioletRed).
		SetTitleAlign(0).
		SetTitleColor(tcell.ColorPaleVioletRed)
	flex.AddItem(form, 30, 1, true)

	return flex

}
