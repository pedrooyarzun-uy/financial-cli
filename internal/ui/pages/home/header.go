package home

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewHeader() *tview.Flex {

	flex := tview.NewFlex()

	left := tview.NewTextView().
		SetText("Welcome Back Pedro 🤩").
		SetTextAlign(tview.AlignLeft)

	right := tview.NewTextView().
		SetText("Cash: $62450").
		SetTextColor(tcell.ColorLimeGreen).
		SetTextAlign(tview.AlignRight)

	flex.SetDirection(tview.FlexColumn)
	flex.AddItem(left, 0, 1, false)
	flex.AddItem(right, 0, 1, false)

	return flex
}
