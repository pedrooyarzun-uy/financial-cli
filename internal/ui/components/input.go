package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewInput(label string, width int, color tcell.Color, background tcell.Color, mask bool) *tview.InputField {
	input := tview.NewInputField().
		SetLabel(label).
		SetFieldWidth(width)

	if mask {
		input.SetMaskCharacter('*')
	}

	addColors(input, color)

	return input
}

func addColors(input *tview.InputField, color tcell.Color) {
	input.SetLabelColor(color).
		SetFieldTextColor(color).
		SetFieldBackgroundColor(tcell.ColorBlack)
}
