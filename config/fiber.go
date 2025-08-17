package config

import (
	"cobaaja/contact-management/app/middleware"
	"errors"
	"os"
	"time"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"
)

func NewFiber(logger *logrus.Logger) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      os.Getenv("APP_NAME"),
		ErrorHandler: ErrorHandlerFiber(),
		BodyLimit:    25 * 1024 * 1024, // 25MB
		// Prefork:      true,
		Concurrency:  256 * 1024, // Atur jumlah worker untuk menangani request
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	})

	// Aktifkan kompresi untuk mengurangi ukuran respons
	app.Use(compress.New())

	// cors middleware
	app.Use(cors.New(cors.ConfigDefault))

	// middleware log method dan path url
	app.Use(middleware.LogRequestMiddleware(logger))

	cfgSwagger := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/contact_management.openapi.json",
		Path:     "docs",
		Title:    "Contact Management API Docs",
	}

	// middleware swagger
	app.Use(swagger.New(cfgSwagger))

	return app
}

func ErrorHandlerFiber() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError
		message := "Internal Server Error"

		// Retrieve the custom status code if it's a *fiber.Error
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
			message = e.Message
		}

		// Set Content-Type: text/plain; charset=utf-8
		// ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

		// Return status code with error message
		return ctx.Status(code).JSON(fiber.Map{
			// "message": err.Error(),
			"message": message,
		})
	}
}
