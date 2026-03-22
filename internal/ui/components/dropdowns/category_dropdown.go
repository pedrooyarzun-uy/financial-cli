package dropdowns

import (
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/services"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"
)

type CategoryDropdown struct {
	*components.DropDown
	categories map[string]int
	service    *services.CategoryService
	onSelected func(categoryId int)
}

func NewCategoryDropdown(label string, labelWidth int, fieldWidth int, onSelected func(categoryID int)) *CategoryDropdown {
	d := components.NewDropDown(label, labelWidth, fieldWidth, []string{}, nil)

	service := services.NewCategoryService(api.CLIENT)

	component := CategoryDropdown{
		DropDown:   d,
		service:    service,
		categories: make(map[string]int),
		onSelected: onSelected,
	}

	component.LoadCategories()

	return &component
}

func (d *CategoryDropdown) LoadCategories() {
	options, err := d.service.GetAllForDropdown("Add new category...", false)

	//If endpoint has an error, load with message
	if err != nil {
		d.AddOption("Something went wrong", nil)
		return
	}

	labels := make([]string, 0, len(options))

	for _, opt := range options {
		labels = append(labels, opt.Label)
		d.categories[opt.Label] = opt.Value
	}

	d.SetOptions(labels, func(text string, index int) {
		categoryId := d.categories[text]

		if d.onSelected != nil {
			d.onSelected(categoryId)
		}
	})
}

func (d *CategoryDropdown) GetSelectedCategoryID() (int, bool) {
	_, label := d.GetCurrentOption()
	categoryID, ok := d.categories[label]
	return categoryID, ok
}
