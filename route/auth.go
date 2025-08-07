package route

import (
	"cobaaja/contact-management/app/controller"
	"cobaaja/contact-management/app/middleware"
	"cobaaja/contact-management/app/repository"
	"cobaaja/contact-management/config"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AuthRoute(app *fiber.App, db *gorm.DB, logger *logrus.Logger, validator *config.Validator) {
	auth := app.Group("/api/v1/auth")

	userRepo := repository.NewUserRepository(db, logger)
	authCtrl := controller.NewAuthController(userRepo, logger, validator)

	auth.Post("/register", authCtrl.Register)
	auth.Post("/login", authCtrl.Login)

	// gunakan jwt middleware
	authPrivate := auth.Use(middleware.JWTMiddleware)

	authPrivate.Post("/me", authCtrl.Me)
	authPrivate.Post("/refresh-token", authCtrl.RefreshToken)
}
