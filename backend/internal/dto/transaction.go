package dto

import "github.com/arganaphang/wallet/backend/internal/entity"

type TransactionAddRequest struct {
	Name     string                 `json:"name"     validate:"required"`
	Amount   uint64                 `json:"amount"   validate:"required"`
	Category string                 `json:"category" validate:"required"`
	Type     entity.TransactionType `json:"type"     validate:"oneof=income outcome"`
}

type TransactionAddResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type TransactionUpdateByIDRequest struct {
	Name     string                 `json:"name"     validate:"required"`
	Amount   uint64                 `json:"amount"   validate:"required"`
	Category string                 `json:"category" validate:"required"`
	Type     entity.TransactionType `json:"type"     validate:"oneof=income outcome"`
}

type TransactionUpdateByIDResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type TransactionGetAllResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Data    []entity.Transaction `json:"data"`
}

type TransactionGetByIDResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    *entity.Transaction `json:"data"`
}

type TransactionDeleteByIDResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    *entity.Transaction `json:"data"`
}
