package services

import (
	"net/url"
	"strconv"

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

func (s *TransactionService) GetTransactionsByDetail(from string, to string, category int, subcategory int, page int, limit int) ([]dto.TransactionByDetail, int, error) {
	res := dto.GetTransactionsByDetailRes{}

	params := url.Values{}

	if from != "" {
		params.Add("from", from)
	}

	if to != "" {
		params.Add("to", to)
	}

	if category != 0 {
		params.Add("category", strconv.Itoa(category))
	}

	if subcategory != 0 {
		params.Add("subcategory", strconv.Itoa(subcategory))
	}

	params.Add("page", strconv.Itoa(page))
	params.Add("limit", strconv.Itoa(limit))

	url := "/transaction/get-all-by-detail?" + params.Encode()

	err := s.apiClient.GetMethod(url, &res)

	if err != nil {
		return nil, 0, err
	}

	return res.Transactions, res.TotalPages, nil
}
