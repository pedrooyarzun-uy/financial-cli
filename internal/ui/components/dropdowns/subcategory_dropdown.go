package dropdowns

import (
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/services"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"
)

type SubCategoryDropdown struct {
	*components.DropDown
	subcategories map[string]int
	service       *services.SubcategoryService
	onSelected    func(subcategoryId int)
}

func NewSubCategoryDropdown(label string, labelWidth int, fieldWidth int) *SubCategoryDropdown {

	d := components.NewDropDown(label, labelWidth, fieldWidth, []string{}, nil)

	service := services.NewSubcategoryService(api.CLIENT)

	component := SubCategoryDropdown{
		DropDown:      d,
		service:       service,
		subcategories: make(map[string]int),
	}

	return &component
}

func (d *SubCategoryDropdown) LoadSubCategories(categoryId int) {
	options, err := d.service.GetAllForDropdown(categoryId, "Add new subcategory...", false)

	if err != nil {
		d.AddOption("Something went wrong", nil)
		return
	}

	labels := make([]string, 0, len(options))

	for _, opt := range options {
		labels = append(labels, opt.Label)
	}

	// Update dropdown options
	d.SetOptions(labels, func(text string, index int) {
		subcategoryId := d.subcategories[text]

		if d.onSelected != nil {
			d.onSelected(subcategoryId)
		}
	})
}
