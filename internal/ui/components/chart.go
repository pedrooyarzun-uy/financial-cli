package components

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
)

type Chart struct {
	*tvxwidgets.BarChart
	labels []string
}

func NewChart() *Chart {
	bc := tvxwidgets.NewBarChart()
	chart := &Chart{
		BarChart: bc,
		labels:   []string{},
	}

	return chart
}

func (c *Chart) Add(label string, value int, color string) {
	var b strings.Builder
	fmt.Fprintf(&b, "[%s]%s[-]: $%d\n", parseHexColor(color), label, value)
	c.BarChart.AddBar(label, value, parseHexColor(color))
	c.labels = append(c.labels, label)
}

func (c *Chart) Reset() {
	c.BarChart.SetMaxValue(0)

	for _, v := range c.labels {
		c.BarChart.RemoveBar(v)
	}
	c.labels = []string{}
}

func parseHexColor(hex string) tcell.Color {
	color := strings.TrimPrefix(hex, "#")
	v, _ := strconv.ParseInt(color, 16, 32)

	return tcell.NewHexColor(int32(v))
}
