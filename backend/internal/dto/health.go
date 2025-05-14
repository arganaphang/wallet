package dto

type HealthzResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
