package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// daftarin route2nya disini
func LoadRoutes(app *fiber.App, db *gorm.DB, logger *logrus.Logger) {
	AuthRoute(app, db, logger)
	ContactRoute(app, db, logger)
}
