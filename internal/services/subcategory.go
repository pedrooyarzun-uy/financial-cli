package services

import (
	"strconv"

	"github.com/pedrooyarzun-uy/financial-cli/internal/api"
	"github.com/pedrooyarzun-uy/financial-cli/internal/dto"
)

type SubcategoryService struct {
	apiClient *api.ApiClient
}

func NewSubcategoryService(client *api.ApiClient) *SubcategoryService {
	return &SubcategoryService{apiClient: client}
}

func (s *SubcategoryService) GetAllForDropdown(id int) ([]dto.DropdownOption, error) {
	response := dto.GetAllSubcategoriesRes{}

	url := "/subcategory/get-by-category?id=" + strconv.Itoa(id)
	err := s.apiClient.GetMethod(url, &response)

	if err != nil {
		return nil, err
	}

	result := make([]dto.DropdownOption, 0, len(response.Subcategories)+1)

	for _, v := range response.Subcategories {
		result = append(result, dto.DropdownOption{
			Label: v.Name,
			Value: v.Id,
		})
	}

	result = append(result, dto.DropdownOption{
		Label: "Add new subcategory...",
		Value: -1,
	})

	return result, nil
}
