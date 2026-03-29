package dto

import "github.com/pedrooyarzun-uy/financial-cli/internal/domain"

type GetAllAcounts struct {
	Message  string `json:"message"`
	Accounts []domain.Account
}

type AddAccountReq struct {
	Name     string `json:"name"`
	Number   string `json:"number"`
	Currency int    `json:"currency"`
	Bank     int    `json:"bank"`
}
