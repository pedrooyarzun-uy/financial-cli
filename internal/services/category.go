package services

import (
	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/dto"
)

type CategoryService struct {
	apiClient *api.ApiClient
}

func NewCategoryService(client *api.ApiClient) *CategoryService {
	return &CategoryService{apiClient: client}
}

func (s *CategoryService) GetAllForDropdown() ([]string, error) {
	response := dto.GetAllCategoriesRes{}

	err := s.apiClient.GetMethod("/category/get-all", &response)

	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(response.Categories)+1)

	for _, v := range response.Categories {
		result = append(result, v.Name)
	}

	result = append(result, "Add new category...")

	return result, nil
}
