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

func ContactRoute(app *fiber.App, db *gorm.DB, logger *logrus.Logger, validator *config.Validator) {
	// init contact repository
	contactRepo := repository.NewContactRepository(db, logger)

	// init contact controller
	contactCtrl := controller.NewContactController(contactRepo, logger, validator)

	// group route
	contactRoute := app.Group("/api/v1/contacts")

	// add jwt middleware
	contactPrivate := contactRoute.Use(middleware.JWTMiddleware)

	contactPrivate.Get("/", contactCtrl.GetAllContact)
	contactPrivate.Get("/:id", contactCtrl.FindContactByID)
	contactPrivate.Post("/", contactCtrl.CreateContact)
	contactPrivate.Put("/:id", contactCtrl.UpdateContactByID)
	contactPrivate.Delete("/:id", contactCtrl.DeleteContactByID)
}
