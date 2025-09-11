package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewInput(label string, width int, color tcell.Color, background tcell.Color) *tview.InputField {
	return tview.NewInputField().
		SetLabel(label).
		SetFieldWidth(width).
		SetLabelColor(color).
		SetFieldTextColor(color).
		SetFieldBackgroundColor(background)

}
