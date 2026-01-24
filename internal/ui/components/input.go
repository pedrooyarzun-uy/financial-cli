package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewInputField(label string, labelWidth int, fieldWidth int) *tview.InputField {
	input := tview.NewInputField().
		SetLabel(label).
		SetLabelWidth(labelWidth).
		SetFieldWidth(fieldWidth).
		SetFieldBackgroundColor(tcell.ColorGray).
		SetFieldTextColor(tcell.ColorWhite).
		SetLabelColor(tcell.ColorWhite)

	return input
}
