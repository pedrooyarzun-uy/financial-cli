package services

import (
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/dto"
)

type TransactionService struct {
	apiClient *api.ApiClient
}

func NewTransactionService(client *api.ApiClient) *TransactionService {
	return &TransactionService{apiClient: client}
}

func (s *TransactionService) Add(amount float64, account int, currency int, type_ int, category int, notes string) error {
	req := dto.AddTransactionReq{
		Amount:   amount,
		Account:  account,
		Currency: currency,
		Type:     type_,
		Category: category,
		Notes:    notes,
	}

	res := dto.AddTransactionRes{}

	err := s.apiClient.PostMethod("/transaction/add", &res, req, false)

	return err

}

func (s *TransactionService) GetTotalsByCategory() ([]dto.TotalByCategory, error) {
	res := dto.TotalsByCategoryRes{}

	err := s.apiClient.GetMethod("/transaction/get-totals-by-category", &res)

	if err != nil {
		return nil, err
	}

	return res.Totals, err
}
