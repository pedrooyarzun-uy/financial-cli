package dropdowns

import (
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/services"
	"github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"
)

type AccountDropdown struct {
	*components.DropDown
	accounts   map[string]int
	service    *services.AccountService
	onSelected func(accountId int)
}

func NewAccountDropdown(label string, labelWidth int, fieldWidth int) *AccountDropdown {

	d := components.NewDropDown(label, labelWidth, fieldWidth, []string{}, nil)

	service := services.NewAccountService(api.CLIENT)

	component := AccountDropdown{
		DropDown: d,
		service:  service,
		accounts: make(map[string]int),
	}

	component.LoadAccounts()

	return &component
}

func (d *AccountDropdown) LoadAccounts() {
	options, err := d.service.GetAllForDropdown()

	if err != nil {
		d.AddOption("Something went wrong", nil)
		return
	}

	labels := make([]string, 0, len(options))

	for _, opt := range options {
		labels = append(labels, opt.Label)
		d.accounts[opt.Label] = opt.Value
	}

	d.SetOptions(labels, nil)
}

func (d *AccountDropdown) GetSelectedAccountID() (int, bool) {
	_, label := d.GetCurrentOption()
	accountID, ok := d.accounts[label]
	return accountID, ok
}
