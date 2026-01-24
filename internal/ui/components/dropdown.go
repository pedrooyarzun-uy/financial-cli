package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewDropDown(label string, labelWidth int, fieldWidth int, options []string, selectedFunc func(text string, index int)) *tview.DropDown {
	dropdown := tview.NewDropDown().
		SetLabel(label).
		SetLabelWidth(labelWidth).
		SetFieldWidth(fieldWidth).
		SetOptions(options, selectedFunc).
		SetFieldBackgroundColor(tcell.ColorGray).
		SetFieldTextColor(tcell.ColorWhite).
		SetLabelColor(tcell.ColorWhite).
		SetListStyles(
			tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack),
			tcell.StyleDefault.Background(tcell.ColorGray).Foreground(tcell.ColorWhite),
		).
		SetPrefixStyle(tcell.StyleDefault.Foreground(tcell.ColorGray).Foreground(tcell.ColorWhite)).
		SetFocusedStyle(tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack))

	return dropdown
}
