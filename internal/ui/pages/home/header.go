package home

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewHeader() *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	topRow := tview.NewFlex().SetDirection(tview.FlexColumn)
	left := tview.NewTextView().
		SetText("Welcome Back Pedro 🤩").
		SetTextAlign(tview.AlignLeft)

	right := tview.NewTextView().
		SetText("Cash: $62450").
		SetTextColor(tcell.ColorLimeGreen).
		SetTextAlign(tview.AlignRight)

	topRow.AddItem(left, 0, 1, false)
	topRow.AddItem(right, 0, 1, false)

	bottomRow := tview.NewTextView().
		SetText("Total spent: $15672").
		SetTextColor(tcell.ColorRed).
		SetTextAlign(tview.AlignRight)

	// Agregamos filas al Flex principal
	flex.AddItem(topRow, 1, 0, false)
	flex.AddItem(bottomRow, 1, 0, false)

	return flex
}
