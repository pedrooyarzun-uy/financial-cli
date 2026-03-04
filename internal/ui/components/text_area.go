package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TextArea struct {
	*tview.TextArea
}

func NewTextArea(label string, placeholder string, maxLength int) *TextArea {
	t := tview.NewTextArea().
		SetLabel(label).
		SetPlaceholder(placeholder).
		SetMaxLength(maxLength)

	t.SetBackgroundColor(tcell.ColorGray)

	textArea := &TextArea{
		TextArea: t,
	}

	return textArea
}
