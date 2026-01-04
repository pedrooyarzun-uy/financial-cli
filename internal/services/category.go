package services

import (
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/dto"
)

type CategoryService struct {
	apiClient *api.ApiClient
}

type DropdownOption struct {
	Label string
	Value int
}

func NewCategoryService(client *api.ApiClient) *CategoryService {
	return &CategoryService{apiClient: client}
}

func (s *CategoryService) GetAllForDropdown() ([]DropdownOption, error) {
	response := dto.GetAllCategoriesRes{}

	err := s.apiClient.GetMethod("/category/get-all", &response)
	if err != nil {
		return nil, err
	}

	result := make([]DropdownOption, 0, len(response.Categories)+1)

	for _, v := range response.Categories {
		result = append(result, DropdownOption{
			Label: v.Name,
			Value: v.Id,
		})
	}

	result = append(result, DropdownOption{
		Label: "Add new category...",
		Value: -1,
	})

	return result, nil
}
