package dto

import "github.com/pedrooyarzun-uy/financial-cli/internal/domain"

type GetAllCurrencies struct {
	Message    string            `json:"message"`
	Currencies []domain.Currency `json:"currencies"`
}
