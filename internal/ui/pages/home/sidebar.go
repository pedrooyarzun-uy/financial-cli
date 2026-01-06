package home

import (
	"fmt"
	"strings"

	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/services"
	"github.com/rivo/tview"
)

func NewSidebar(app *tview.Application, pages *tview.Pages) *tview.Flex {

	ts := services.NewTransactionService(api.CLIENT)

	sidebar := tview.NewFlex()

	categories, err := ts.GetTotalsByCategory()

	list := tview.NewTextView().SetDynamicColors(true)

	if err != nil {
		list.SetText(err.Error())
		sidebar.AddItem(list, 0, 1, true)
		return sidebar
	}

	var b strings.Builder
	for _, c := range categories {
		fmt.Fprintf(&b, "[%s]%s[-]: $%.2f\n", c.Color, c.Category, c.Total)
	}

	list.SetText(b.String())

	sidebar.AddItem(list, 0, 1, true)

	return sidebar
}
