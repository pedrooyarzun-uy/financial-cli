package dropdowns

import "github.com/pedrooyarzun-uy/financial-cli/internal/ui/components"

type BankDropdown struct {
	*components.DropDown
	banks map[string]int
}

func NewBankDropdown(label string, labelWidth int, fieldWidth int) *BankDropdown {
	d := components.NewDropDown(label, labelWidth, fieldWidth, []string{"Itaú", "Scotiabank", "Banco república", "Santander", "BBVA", "Prex", "Mi Dinero"}, nil)

	component := BankDropdown{
		DropDown: d,
		banks: map[string]int{
			"Itaú":            1,
			"Scotiabank":      2,
			"Banco república": 3,
			"Santander":       4,
			"BBVA":            5,
			"Prex":            6,
			"Mi Dinero":       7,
		},
	}

	return &component
}

func (d *BankDropdown) GetSelectedBankID() (int, bool) {
	_, label := d.GetCurrentOption()
	bankID, ok := d.banks[label]
	return bankID, ok
}
