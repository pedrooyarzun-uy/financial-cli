package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewButton(label string, selectedFunc func()) *tview.Button {
	button := tview.NewButton(label).
		SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorGray)).
		SetSelectedFunc(selectedFunc)

	return button
}
