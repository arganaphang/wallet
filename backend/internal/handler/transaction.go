package handler

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	ulid "github.com/oklog/ulid/v2"
	"github.com/sirupsen/logrus"

	"github.com/arganaphang/wallet/backend/internal/dto"
	"github.com/arganaphang/wallet/backend/internal/entity"
	"github.com/arganaphang/wallet/backend/internal/service"
)

type ITransactionHandler interface {
	Add(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	UpdateByID(ctx *fiber.Ctx) error
	DeleteByID(ctx *fiber.Ctx) error
}

type transactionHandler struct {
	services service.Service
}

func NewTransactionHandler(app *fiber.App, services service.Service) ITransactionHandler {
	handler := &transactionHandler{services: services}

	route := app.Group("/api/v1/transactions")
	route.Post("/", handler.Add)
	route.Get("/", handler.GetAll)
	route.Get("/:id", handler.GetByID)
	route.Put("/:id", handler.UpdateByID)
	route.Delete("/:id", handler.DeleteByID)

	return handler
}

// Add implements ITransactionHandler.
func (t *transactionHandler) Add(ctx *fiber.Ctx) error {
	var data dto.TransactionAddRequest
	if err := ctx.BodyParser(&data); err != nil {
		logrus.Info("Error ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(dto.TransactionAddResponse{
			Success: false,
			Message: "failed to parse request body",
		})
	}

	if err := t.services.Transaction.Add(ctx.Context(), entity.Transaction{
		Name:     data.Name,
		Amount:   data.Amount,
		Category: data.Category,
		Type:     data.Type,
	}); err != nil {
		logrus.Info("Error ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(dto.TransactionAddResponse{
			Success: false,
			Message: "failed to add transaction",
		})
	}

	return ctx.Status(http.StatusCreated).JSON(dto.TransactionAddResponse{
		Success: true,
		Message: "transaction added",
	})
}

// DeleteByID implements ITransactionHandler.
func (t *transactionHandler) DeleteByID(ctx *fiber.Ctx) error {
	id, err := ulid.Parse(ctx.Params("id"))
	if err != nil {
		logrus.Info("Error ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(dto.TransactionDeleteByIDResponse{
			Success: false,
			Message: "failed to parse id transaction",
		})
	}

	if err := t.services.Transaction.DeleteByID(ctx.Context(), id); err != nil {
		logrus.Info("Error ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(dto.TransactionDeleteByIDResponse{
			Success: false,
			Message: "failed to delete transaction",
		})
	}

	return ctx.Status(http.StatusOK).JSON(dto.TransactionDeleteByIDResponse{
		Success: true,
		Message: "transaction deleted",
	})
}

// GetTransactionByID implements ITransactionHandler.
func (t *transactionHandler) GetByID(ctx *fiber.Ctx) error {
	id, err := ulid.Parse(ctx.Params("id"))
	if err != nil {
		logrus.Info("Error ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(dto.TransactionGetByIDResponse{
			Success: false,
			Message: "failed to parse id transaction",
		})
	}
	result, err := t.services.Transaction.GetByID(ctx.Context(), id)
	if err != nil {
		logrus.Info("Error ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(dto.TransactionGetByIDResponse{
			Success: false,
			Message: "failed to get transaction by id",
		})
	}
	return ctx.Status(http.StatusOK).JSON(dto.TransactionGetByIDResponse{
		Success: true,
		Message: "get transaction by id success",
		Data:    result,
	})
}

// GetTransactions implements ITransactionHandler.
func (t *transactionHandler) GetAll(ctx *fiber.Ctx) error {
	result, err := t.services.Transaction.GetAll(ctx.Context())
	if err != nil {
		logrus.Info("Error ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(dto.TransactionGetAllResponse{
			Success: false,
			Message: "failed to get transactions",
		})
	}
	return ctx.Status(http.StatusOK).JSON(dto.TransactionGetAllResponse{
		Success: true,
		Message: "get transactions success",
		Data:    result,
	})
}

// UpdateByID implements ITransactionHandler.
func (t *transactionHandler) UpdateByID(ctx *fiber.Ctx) error {
	id, err := ulid.Parse(ctx.Params("id"))
	if err != nil {
		logrus.Info("Error ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(dto.TransactionUpdateByIDResponse{
			Success: false,
			Message: "failed to parse id transaction",
		})
	}

	var data dto.TransactionUpdateByIDRequest
	if err := ctx.BodyParser(&data); err != nil {
		logrus.Info("Error ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(dto.TransactionUpdateByIDResponse{
			Success: false,
			Message: "failed to parse request body",
		})
	}

	if err := t.services.Transaction.UpdateByID(ctx.Context(), id, entity.Transaction{
		Name:     data.Name,
		Amount:   data.Amount,
		Category: data.Category,
		Type:     data.Type,
	}); err != nil {
		logrus.Info("Error ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(dto.TransactionUpdateByIDResponse{
			Success: false,
			Message: "failed to update transaction",
		})
	}

	return ctx.Status(http.StatusOK).JSON(dto.TransactionUpdateByIDResponse{
		Success: true,
		Message: "transaction updated",
	})
}
