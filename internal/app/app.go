package app

import (
	"context"
	"fmt"

	"em/internal/http/handlers"
	"em/internal/storage/postgre"
	"em/pkg/enricher"

	user "em/internal/http/handlers/user"
	srvuser "em/internal/services/user"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type App struct {
	*fiber.App
}

func New(ctx context.Context, conn string) (*App, error) {
	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.ErrorHanlder,
	})

	app.Use(swagger.New(swagger.Config{
		BasePath: "/api/v1",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
	}))

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(logger.New())

	log.Info("start init postgre")

	storage, err := postgre.New(context.Background(), conn)
	if err != nil {
		return nil, fmt.Errorf("postgre.New: %v", err)
	}

	log.Info("start init services")

	userService := srvuser.New(storage, enricher.New())

	v1 := app.Group("/api/v1")

	user.Register(v1, userService, storage)

	return &App{app}, nil
}
