package dto

import "github.com/pedrooyarzun-uy/financial-cli/internal/domain"

type GetAllSubcategoriesRes struct {
	Message       string               `json:"message" binding:"required"`
	Subcategories []domain.Subcategory `json:"subcategories" binding:"required"`
}
