package main

import (
	"log"
	"net/http"
	"os"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	_ "github.com/arganaphang/wallet/backend/docs"
	"github.com/arganaphang/wallet/backend/internal/handler"
	"github.com/arganaphang/wallet/backend/internal/repository"
	"github.com/arganaphang/wallet/backend/internal/service"
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
	app.Get("/healthz", getHealthz)
	app.Get("/swagger/*", swagger.HandlerDefault)

	if err := app.Listen("0.0.0.0:8000"); err != nil {
		log.Fatalln(err)
	}
}

type HealthzResponse struct {
	Message string `json:"message"`
}

// @Description Health Check
// @ID healthz
// @Produce json
// @Success 200 {object} HealthzResponse "OK"
// @Router /healthz [get]
func getHealthz(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(HealthzResponse{
		Message: "OK",
	})
}
