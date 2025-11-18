package dto

import "github.com/arganaphang/wallet/backend/internal/entity"

type CategoryAddRequest struct {
	Name string `json:"name"`
}

type CategoryAddResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type CategoryDeleteByIDResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type CategoryGetAllResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    []entity.Category `json:"data"`
}
