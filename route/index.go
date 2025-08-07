package route

import (
	"cobaaja/contact-management/config"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// daftarin route2nya disini
func LoadRoutes(app *fiber.App, db *gorm.DB, logger *logrus.Logger, validator *config.Validator) {
	AuthRoute(app, db, logger, validator)
	ContactRoute(app, db, logger, validator)
}
