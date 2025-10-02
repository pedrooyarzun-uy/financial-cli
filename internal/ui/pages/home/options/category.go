package options

import (
	"github.com/gdamore/tcell/v2"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"
	"github.com/rivo/tview"
)

func NewCategory(pages *tview.Pages) tview.Primitive {

	form := tview.NewForm().
		AddInputField("Name:", "", 30, nil, nil).
		SetLabelColor(tcell.ColorCornflowerBlue).
		SetFieldBackgroundColor(tcell.ColorDarkSlateGrey).
		AddButton("OK", func() {
			pages.RemovePage("category")
		}).
		AddButton("Cancel", func() {
			pages.RemovePage("category")
		})

	form.SetBorder(true).
		SetTitle("NEW CATEGORY").
		SetTitleColor(tcell.ColorYellow).
		SetTitleAlign(tview.AlignCenter).
		SetBackgroundColor(tcell.ColorLightSlateGray)

	modal := components.NewFormModal(pages, form)

	return modal
}
