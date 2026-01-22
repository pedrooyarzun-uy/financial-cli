package components

import "github.com/rivo/tview"

func NewDropDown(label string, labelWidth int, fieldWidth int, options []string, selectedFunc func(text string, index int)) *tview.DropDown {
	dropdown := tview.NewDropDown().
		SetLabel(label).
		SetLabelWidth(labelWidth).
		SetFieldWidth(fieldWidth).
		SetOptions(options, selectedFunc)

	return dropdown
}
