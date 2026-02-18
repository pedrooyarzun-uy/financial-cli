package components

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
)

func NewChart() *tvxwidgets.BarChart {
	chart := tvxwidgets.NewBarChart()

	return chart
}

func AddItem(chart *tvxwidgets.BarChart, label string, value int, color string) {
	var b strings.Builder
	fmt.Fprintf(&b, "[%s]%s[-]: $%d\n", parseHexColor(color), label, value)
	chart.AddBar(label[:7], value, parseHexColor(color))
}

func parseHexColor(hex string) tcell.Color {
	color := strings.TrimPrefix(hex, "#")
	v, _ := strconv.ParseInt(color, 16, 32)

	return tcell.NewHexColor(int32(v))
}
