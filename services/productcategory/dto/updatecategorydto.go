package dto

import (
	"backend_capstone/models"

	"github.com/google/uuid"
)

type UpdateCategoryDTO struct {
	Name        string `validator:"required"`
	IsAvailable bool   `validator:"required,boolean"`
	Description string `validator:"required,alphaunicode"`
}

func (data *UpdateCategoryDTO) GenerateModel(id uuid.UUID, slug string) *models.ProductCategory {
	return &models.ProductCategory{
		Id:          id,
		Slug:        slug,
		Name:        data.Name,
		IsAvailable: data.IsAvailable,
		Description: data.Description,
	}
}
