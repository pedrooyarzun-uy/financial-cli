package dto

import "github.com/pedrooyarzun-uy/financial-cli/internal/domain"

type GetAllCategoriesRes struct {
	Message    string            `json:"message"`
	Categories []domain.Category `json:"categories"`
}
