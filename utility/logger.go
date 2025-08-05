package utility

import (
	"cobaaja/contact-management/config"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// Fungsi utilitas untuk mendapatkan logger dari konteks Fiber
func GetLogger(ctx *fiber.Ctx) *logrus.Entry {
	if l, ok := ctx.Locals("logger").(*logrus.Entry); ok {
		return l
	}
	// Fallback ke logger default jika tidak ada di konteks
	return config.NewLogger().WithFields(logrus.Fields{})
}
