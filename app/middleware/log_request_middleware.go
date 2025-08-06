package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func LogRequestMiddleware(logger *logrus.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger.Info(fmt.Sprintf("%s %s", ctx.Method(), ctx.Path()))

		return ctx.Next()
	}
}
