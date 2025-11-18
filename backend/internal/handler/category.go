package handler

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	ulid "github.com/oklog/ulid/v2"
	"github.com/sirupsen/logrus"

	"github.com/arganaphang/wallet/backend/internal/dto"
	"github.com/arganaphang/wallet/backend/internal/service"
)

type ICategoryHandler interface {
	Add(ctx *fiber.Ctx) error
	DeleteByID(ctx *fiber.Ctx) error
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
	route.Delete("/:id", handler.DeleteByID)
	return handler
}

// Add implements ICategoryHandler.
// @Summary Create
// @Description Create Category
// @ID category-create
// @Tags Category
// Accept json
// @Produce json
// @Success 200 {object} dto.CategoryAddResponse "OK"
// @Router /api/v1/categories [post]
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
// @Summary Get All
// @Description Get All Category
// @ID category-get-all
// @Tags Category
// Accept json
// @Produce json
// @Success 200 {object} dto.CategoryGetAllResponse "OK"
// @Router /api/v1/categories [get]
func (c *categoryHandler) GetAll(ctx *fiber.Ctx) error {
	results, err := c.services.Category.GetAll(ctx.Context(), ctx.Query("q"))
	if err != nil {
		logrus.Info("Error ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(dto.CategoryGetAllResponse{
			Success: false,
			Message: "failed to get category",
		})
	}

	return ctx.Status(http.StatusOK).JSON(dto.CategoryGetAllResponse{
		Success: true,
		Message: "get category",
		Data:    results,
	})
}

// DeleteByID implements ICategoryHandler.
// @Summary Delete by ID
// @Description Delete category by id
// @ID category-delete-by-id
// @Tags Category
// Accept json
// @Produce json
// @Success 200 {object} dto.CategoryDeleteByIDResponse "OK"
// @Router /api/v1/categories/:id [delete]
func (c *categoryHandler) DeleteByID(ctx *fiber.Ctx) error {
	id, err := ulid.Parse(ctx.Params("id"))
	if err != nil {
		logrus.Info("Error ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(dto.CategoryDeleteByIDResponse{
			Success: false,
			Message: "failed to parse id transaction",
		})
	}
	if err := c.services.Category.DeleteByID(ctx.Context(), id); err != nil {
		logrus.Info("Error ", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(dto.CategoryDeleteByIDResponse{
			Success: false,
			Message: "failed to get category",
		})
	}

	return ctx.Status(http.StatusOK).JSON(dto.CategoryDeleteByIDResponse{
		Success: true,
		Message: "category deleted",
	})
}
