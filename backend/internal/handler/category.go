package handler

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/arganaphang/wallet/backend/internal/dto"
	"github.com/arganaphang/wallet/backend/internal/service"
)

type ICategoryHandler interface {
	Add(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
}

type categoryHandler struct {
	services service.Service
}

func NewCategoryHandler(app *fiber.App, services service.Service) ICategoryHandler {
	handler := &categoryHandler{services: services}

	route := app.Group("/api/v1/categories")
	route.Get("/", handler.GetAll)
	route.Post("/", handler.Add)
	return handler
}

// Add implements ICategoryHandler.
func (c *categoryHandler) Add(ctx *fiber.Ctx) error {
	var data dto.CategoryAddRequest
	if err := ctx.BodyParser(&data); err != nil {
		logrus.Info("Error ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(dto.CategoryAddResponse{
			Success: false,
			Message: "failed to parse request body",
		})
	}

	if err := c.services.Category.Add(ctx.Context(), data.Name); err != nil {
		logrus.Info("Error ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(dto.CategoryAddResponse{
			Success: false,
			Message: "failed to add category",
		})
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CategoryAddResponse{
		Success: true,
		Message: "category added",
	})
}

// GetAll implements ICategoryHandler.
func (c *categoryHandler) GetAll(ctx *fiber.Ctx) error {
	results, err := c.services.Category.GetAll(ctx.Context(), ctx.Query("q"))
	if err != nil {
		logrus.Info("Error ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(dto.CategoryGetAllResponse{
			Success: false,
			Message: "failed to get category",
		})
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CategoryGetAllResponse{
		Success: true,
		Message: "get category success",
		Data:    results,
	})
}
