package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Form struct {
	*tview.Form
}

func NewForm() *Form {
	f := tview.NewForm().
		SetFieldBackgroundColor(tcell.ColorGray).
		SetFieldTextColor(tcell.ColorWhite).
		SetLabelColor(tcell.ColorWhite)

	form := &Form{
		Form: f,
	}

	return form
}
