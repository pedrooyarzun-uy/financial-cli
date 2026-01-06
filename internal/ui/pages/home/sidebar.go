package home

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/services"
	"github.com/rivo/tview"
)

func NewSidebar(app *tview.Application, pages *tview.Pages) *tview.Flex {

	ts := services.NewTransactionService(api.CLIENT)
	sidebar := tview.NewFlex().SetDirection(tview.FlexRow)
	categories, err := ts.GetTotalsByCategory()

	//Graphic items
	list := tview.NewTextView().SetDynamicColors(true)
	bc := tvxwidgets.NewBarChart()

	bc.SetTitle("Gastos por categoría")
	bc.SetBorder(false)

	if err != nil {
		list.SetText(err.Error())
		sidebar.AddItem(list, 0, 1, true)
		return sidebar
	}

	var b strings.Builder
	for _, c := range categories {
		fmt.Fprintf(&b, "[%s]%s[-]: $%.2f\n", c.Color, c.Category, c.Total)

		c.Color = strings.TrimPrefix(c.Color, "#")
		v, _ := strconv.ParseInt(c.Color, 16, 32)

		bc.AddBar(c.Category[:7], int(c.Total), tcell.NewHexColor(int32(v)))
	}

	list.SetText(b.String())

	sidebar.AddItem(list, 0, 1, true)
	sidebar.AddItem(bc, 0, 1, true)

	return sidebar
}
