package route

import (
	"cobaaja/contact-management/app/controller"
	"cobaaja/contact-management/app/middleware"
	"cobaaja/contact-management/app/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRoute(app *fiber.App, db *gorm.DB) {
	auth := app.Group("/api/v1/auth")

	userRepo := repository.NewUserRepository(db)
	authCtrl := controller.NewAuthController(userRepo)

	auth.Post("/register", authCtrl.Register)
	auth.Post("/login", authCtrl.Login)

	// gunakan jwt middleware
	authPrivate := auth.Use(middleware.JWTMiddleware)

	authPrivate.Post("/me", authCtrl.Me)
}
