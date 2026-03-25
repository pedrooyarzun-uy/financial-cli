package dropdowns

import (
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/services"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"
)

type CurrencyDropdown struct {
	*components.DropDown
	currencies map[string]int
	service    *services.CurrencyService
	onSelected func(accountId int)
}

func NewCurrencyDropdown(label string, labelWidth int, fieldWidth int) *CurrencyDropdown {

	d := components.NewDropDown(label, labelWidth, fieldWidth, []string{}, nil)

	service := services.NewCurrencyService(api.CLIENT)

	component := CurrencyDropdown{
		DropDown:   d,
		service:    service,
		currencies: make(map[string]int),
	}

	component.LoadAccounts()

	return &component
}

func (d *CurrencyDropdown) LoadAccounts() {
	options, err := d.service.GetAllForDropdown()

	if err != nil {
		d.AddOption("Something went wrong", nil)
		return
	}

	labels := make([]string, 0, len(options))

	for _, opt := range options {
		labels = append(labels, opt.Label)
		d.currencies[opt.Label] = opt.Value
	}

	d.SetOptions(labels, nil)
}

func (d *CurrencyDropdown) GetSelectedCurrencyID() (int, bool) {
	_, label := d.GetCurrentOption()
	accountID, ok := d.currencies[label]
	return accountID, ok
}
