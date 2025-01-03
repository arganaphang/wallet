package main

import (
	"log"
	"net/http"
	"os"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/arganaphang/wallet/backend/internal/handler"
	"github.com/arganaphang/wallet/backend/internal/repository"
	"github.com/arganaphang/wallet/backend/internal/service"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

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

	app.Get("/healthz", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(map[string]string{
			"message": "OK",
		})
	})

	if err := app.Listen("0.0.0.0:8000"); err != nil {
		log.Fatalln(err)
	}
}
