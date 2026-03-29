package dropdowns

import "github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"

type TypeDropdown struct {
	*components.DropDown
	types      map[string]int
	onSelected func(typeId int)
}

func NewTypeDropdown(label string, labelWidth int, fieldWidth int) *TypeDropdown {
	d := components.NewDropDown(label, labelWidth, fieldWidth, []string{}, nil)

	component := TypeDropdown{
		DropDown: d,
		types: map[string]int{
			"Income":     1,
			"Outcome":    2,
			"Adjustment": 3,
		},
	}

	d.SetOptions([]string{"Income", "Outcome", "Adjustment"}, nil)

	return &component

}

func (d *TypeDropdown) GetSelectedTypeID() (int, bool) {
	_, label := d.GetCurrentOption()
	typeID, ok := d.types[label]
	return typeID, ok
}
