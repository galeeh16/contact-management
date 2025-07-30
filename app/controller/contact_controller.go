package controller

import (
	"cobaaja/contact-management/app/dto"
	"cobaaja/contact-management/app/entity"
	"cobaaja/contact-management/app/repository"
	"cobaaja/contact-management/utility"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ContactController struct {
	Repo *repository.ContactRepository
}

// initialize contact controller
func NewContactController(repo *repository.ContactRepository) *ContactController {
	return &ContactController{Repo: repo}
}

func (ctrl *ContactController) GetAllContact(ctx *fiber.Ctx) error {
	return ctx.JSON("Not Implemented Yet.")
}

func (ctrl *ContactController) CreateContact(ctx *fiber.Ctx) error {
	req := new(dto.CreateContactRequest)

	// bind request into struct
	ctx.BodyParser(&req)

	// validasi request struct
	v := utility.NewValidator()

	// register custom validasi unique contact phone number
	v.Validate.RegisterValidation("unique_contact_phone", func(fl validator.FieldLevel) bool {
		existsPhone := ctrl.Repo.CheckExistsPhone(req.Phone)
		return !existsPhone
	})

	arrayError := v.ValidateStruct(req)

	if arrayError != nil {
		return utility.BadRequestResponse("Invalid Data", arrayError, ctx)
	}

	// ambil user_id dari ctx local yang dibuat di jwt middleware
	userId, ok := ctx.Locals("user_id").(uint64)
	if !ok {
		return utility.BadRequestResponse("User ID tidak valid", nil, ctx)
	}

	// mapping request struct into entity contact struct
	data := &entity.Contact{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		UserID:    userId,
	}

	// begin transaction
	tx := ctrl.Repo.DB.Begin()

	// create contact
	contact, err := ctrl.Repo.CreateContact(tx, data)
	if err != nil {
		// rollback transaction
		tx.Rollback()

		fmt.Println(err.Error())
		return utility.ErrorResponse("Internal Server Error", nil, ctx)
	}

	// mapping new contact ke response
	contactResponse := &dto.CreateContactResponse{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Phone:     contact.Phone,
		CreatedAt: contact.CreatedAt,
		UpdatedAt: contact.UpdatedAt,
	}

	// commit transaction
	tx.Commit()

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success Create Contact",
		"data":    contactResponse,
	})
}

func (ctrl *ContactController) FindContactByPhone(ctx *fiber.Ctx) error {
	phone := ctx.Params("phone")

	contact, err := ctrl.Repo.FindContactByPhone(phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utility.BadRequestResponse("Contact "+phone+" tidak ditemukan", nil, ctx)
		} else {
			return utility.ErrorResponse("Internal Server Error", nil, ctx)
		}
	}

	// mapping contact entity ke contact response
	contactRes := &dto.ContactResponse{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Phone:     contact.Phone,
		CreatedAt: contact.CreatedAt,
		UpdatedAt: contact.UpdatedAt,
	}

	return utility.SuccessResponse("Success", contactRes, ctx)
}

func (ctrl *ContactController) UpdateContactByID(ctx *fiber.Ctx) error {
	return ctx.JSON("Not Implemented Yet.")
}

func (ctrl *ContactController) DeleteContactByID(ctx *fiber.Ctx) error {
	return ctx.JSON("Not Implemented Yet.")
}
