package services

import (
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/dto"
)

type AccountService struct {
	apiClient *api.ApiClient
}

func NewAccountService(client *api.ApiClient) *AccountService {
	return &AccountService{apiClient: client}
}

func (s *AccountService) GetAllForDropdown() ([]DropdownOption, error) {
	response := dto.GetAllAcounts{}

	err := s.apiClient.GetMethod("/account/get-all", &response)

	if err != nil {
		return nil, err
	}

	result := make([]DropdownOption, 0, len(response.Accounts)+1)

	for _, v := range response.Accounts {
		result = append(result, DropdownOption{
			Label: v.Name,
			Value: v.Id,
		})
	}

	return result, nil
}
