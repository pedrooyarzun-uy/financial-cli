package dto

import "github.com/pedrooyarzun-uy/financial-cli/internal/domain"

type GetAllCurrenciesRes struct {
	Message    string            `json:"message"`
	Currencies []domain.Currency `json:"currencies"`
}
