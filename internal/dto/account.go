package dto

import "github.com/pedrooyarzun-uy/financial-cli/internal/domain"

type GetAllAcounts struct {
	Message  string `json:"message"`
	Accounts []domain.Account
}
