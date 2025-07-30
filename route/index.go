package route

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// daftarin route2nya disini
func LoadRoutes(app *fiber.App, db *gorm.DB) {
	AuthRoute(app, db)
	ContactRoute(app, db)
}
