package services

import (
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/dto"
)

type CurrencyService struct {
	apiClient *api.ApiClient
}

func NewCurrencyService(client *api.ApiClient) *CurrencyService {
	return &CurrencyService{apiClient: client}
}

func (s *CurrencyService) GetAllForDropdown() ([]DropdownOption, error) {
	response := dto.GetAllCurrenciesRes{}

	err := s.apiClient.GetMethod("/currency/get-all", &response)

	if err != nil {
		return nil, err
	}

	result := make([]DropdownOption, 0, len(response.Currencies)+1)

	for _, v := range response.Currencies {
		result = append(result, DropdownOption{
			Label: v.Code,
			Value: v.Id,
		})
	}

	return result, nil
}
