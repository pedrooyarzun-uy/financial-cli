package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type InputField struct {
	*tview.InputField
}

func NewInputField(label string, labelWidth int, fieldWidth int) *InputField {
	i := tview.NewInputField().
		SetLabel(label).
		SetLabelWidth(labelWidth).
		SetFieldWidth(fieldWidth).
		SetFieldBackgroundColor(tcell.ColorGray).
		SetFieldTextColor(tcell.ColorWhite).
		SetLabelColor(tcell.ColorWhite)

	input := &InputField{
		InputField: i,
	}

	return input
}
