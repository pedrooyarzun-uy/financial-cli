package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Button struct {
	*tview.Button
}

func NewButton(label string, selectedFunc func()) *Button {
	button := tview.NewButton(label).
		SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorGray)).
		SetSelectedFunc(selectedFunc)

	btn := &Button{
		Button: button,
	}

	return btn
}
