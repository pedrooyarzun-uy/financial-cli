package services

import (
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/dto"
)

type CreditCardService struct {
	apiClient *api.ApiClient
}

func NewCreditCardService(client *api.ApiClient) *CreditCardService {
	return &CreditCardService{apiClient: client}
}

func (s *CreditCardService) Add(name string, bankID int, closeDay int, dueDay int, creditLimit float64) error {
	req := dto.AddCreditCardReq{
		Name:        name,
		BankID:      bankID,
		CloseDay:    closeDay,
		DueDay:      dueDay,
		CreditLimit: creditLimit,
	}

	res := dto.AddCreditCardRes{}

	err := s.apiClient.PostMethod("/credit-card/add", &res, req, false)

	return err
}
