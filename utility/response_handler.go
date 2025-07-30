package utility

import (
	"errors"
	"os"

	"github.com/gofiber/fiber/v2"
)

// Default error handler
var ResponseDefaultErrorHandler = func(ctx *fiber.Ctx, err error) error {
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

// buat success response (200)
func SuccessResponse(message string, data any, ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": message,
		"data":    data,
	})
}

// buat bad request response
func BadRequestResponse(message string, data any, ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": message,
		"errors":  data,
	})
}

// buat not found response (404)
func NofFoundResponse(message string, data any, ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": message,
		"errors":  data,
	})
}

// buat status unauthorized (401)
func UnauthenticatedResponse(message string, data any, ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": message,
		"errors":  data,
	})
}

// buat status forbidden (403)
func ForbiddenResponse(message string, data any, ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"message": message,
		"errors":  nil,
	})
}

// ini buat general error (500)
func ErrorResponse(message string, data any, ctx *fiber.Ctx) error {
	appLevel := os.Getenv("APP_LEVEL")
	if appLevel == "production" || appLevel == "prod" {
		data = "Whoops, something went wrong"
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": message,
		"errors":  data,
	})
}
