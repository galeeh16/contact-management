package main

import (
	"cobaaja/contact-management/config"
	"cobaaja/contact-management/route"
	"cobaaja/contact-management/utility"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func main() {
	// Load .env file
	config.LoadEnv()

	// Connect database
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}

	// init logger
	// logger := config.NewLogger()

	// Fiber App Init
	app := fiber.New(fiber.Config{
		AppName:      os.Getenv("APP_NAME"),
		ErrorHandler: utility.ResponseDefaultErrorHandler,
		BodyLimit:    25 * 1024 * 1024, // 25MB
		Prefork:      true,
		Concurrency:  256 * 1024, // Atur jumlah worker untuk menangani request
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	})

	// Aktifkan kompresi untuk mengurangi ukuran respons
	app.Use(compress.New())

	// // middleware trace id
	// app.Use(func(ctx *fiber.Ctx) error {
	// 	traceID := uuid.New().String()

	// 	ctx.Locals("trace_id", traceID)

	// 	reqLogger := logger.WithFields(logrus.Fields{
	// 		"trace_id": traceID,
	// 	})

	// 	ctx.Locals("logger", reqLogger)

	// 	reqLogger.Info(fmt.Sprintf("%s %s", ctx.Method(), ctx.Path()))

	// 	return ctx.Next()
	// })

	// load routes
	route.LoadRoutes(app, db)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
