package components

import "github.com/rivo/tview"

func NewFormModal(pages *tview.Pages, form *tview.Form) tview.Primitive {

	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(form, 10, 1, true).
				AddItem(nil, 0, 1, false),
			60, 1, true).
		AddItem(nil, 0, 1, false)
}
