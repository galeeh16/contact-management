package route

import (
	"cobaaja/contact-management/app/controller"
	"cobaaja/contact-management/app/middleware"
	"cobaaja/contact-management/app/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ContactRoute(app *fiber.App, db *gorm.DB) {
	// init contact repository
	contactRepo := repository.NewContactRepository(db)

	// init contact controller
	contactCtrl := controller.NewContactController(contactRepo)

	// group route
	contactRoute := app.Group("/api/v1/contacts")

	// add jwt middleware
	contactPrivate := contactRoute.Use(middleware.JWTMiddleware)

	contactPrivate.Get("/", contactCtrl.GetAllContact)
	contactPrivate.Get("/:phone", contactCtrl.FindContactByPhone)
	contactPrivate.Post("/", contactCtrl.CreateContact)
	contactPrivate.Put("/:id", contactCtrl.UpdateContactByID)
	contactPrivate.Delete("/:id", contactCtrl.DeleteContactByID)
}
