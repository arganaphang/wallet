package main

import (
	"log"
	"net/http"
	"os"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	_ "github.com/arganaphang/wallet/backend/docs"
	"github.com/arganaphang/wallet/backend/internal/dto"
	"github.com/arganaphang/wallet/backend/internal/handler"
	"github.com/arganaphang/wallet/backend/internal/repository"
	"github.com/arganaphang/wallet/backend/internal/service"
	"github.com/arganaphang/wallet/backend/pkg/scalar"
)

// @title Wallet API
// @version 1.0
// @description This is a server for a wallet API.
// @contact.name Argana Phangquestian
// @contact.email arganaphangquestian@gmail.com
// @license.name MIT License
// @license.url https://mit-license.org/
// @host localhost:8000
// @BasePath /
func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/healthz", getHealthz)
	app.Get("/docs", func(ctx *fiber.Ctx) error {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Bookshelf API",
			},
			DarkMode: true,
		})
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(map[string]any{})
		}

		ctx.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

		return ctx.Status(http.StatusOK).SendString(htmlContent)
	})

	db, err := sqlx.Open(
		"postgres",
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		log.Fatalln("failed to open database connection", err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Fatalln("failed to ping database", err.Error())
	}

	repositories := repository.Repository{
		Transaction: repository.NewTransactionRepository(db),
		Category:    repository.NewCategoryRepository(db),
	}

	services := service.Service{
		Transaction: service.NewTransactionService(repositories),
		Category:    service.NewCategoryService(repositories),
	}

	_ = handler.Handler{
		Transaction: handler.NewTransactionHandler(app, services),
		Category:    handler.NewCategoryHandler(app, services),
	}

	if err := app.Listen("0.0.0.0:8000"); err != nil {
		log.Fatalln(err)
	}
}

type HealthzResponse struct {
	Message string `json:"message"`
}

// @Summary Health
// @Description Health Check
// @ID healthz
// @Tags Health
// @Produce json
// @Success 200 {object} HealthzResponse "OK"
// @Router /healthz [get]
func getHealthz(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(dto.HealthzResponse{
		Success: true,
		Message: "OK",
	})
}
