package controller

import (
	"cobaaja/contact-management/app/dto"
	"cobaaja/contact-management/app/entity"
	"cobaaja/contact-management/app/repository"
	"cobaaja/contact-management/utility"
	"errors"
	"fmt"
	"strconv"

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
	// Ambil query params
	pageStr := ctx.Query("page", "1")
	sizeStr := ctx.Query("size", "10")

	// Konversi ke integer
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil || size < 1 {
		size = 10
	}

	// ambil data pagination dari repository
	contacts, total, err := ctrl.Repo.GetAllContact(page, size)
	if err != nil {
		return utility.ErrorResponse("Internal Server Error", nil, ctx)
	}

	// Mapping Contact ke Contacts' DTO
	contactsDTO := make([]dto.ContactResponse, len(contacts))

	for i, contact := range contacts {
		contactsDTO[i] = dto.ContactResponse{
			ID:        contact.ID,
			FirstName: contact.FirstName,
			LastName:  contact.LastName,
			Phone:     contact.Phone,
			CreatedAt: contact.CreatedAt,
			UpdatedAt: contact.UpdatedAt,
		}
	}

	// Hitung total halaman
	totalPages := int((total + int64(size) - 1) / int64(size))

	return utility.SuccessResponse("Success Get Data", fiber.Map{
		"items":       contactsDTO,
		"total":       total,
		"page":        page,
		"size":        size,
		"total_pages": totalPages,
	}, ctx)
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
	id := ctx.Params("id")
	idInt, _ := strconv.ParseUint(id, 10, 64)

	// find contact by id
	_, err := ctrl.Repo.FindContactByID(idInt)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utility.BadRequestResponse("Contact "+id+"Not Found", nil, ctx)
		} else {
			fmt.Println(err.Error())
			return utility.ErrorResponse("Internal Server Error", nil, ctx)
		}
	}

	// delete contact by id
	err = ctrl.Repo.DeleteContactByID(idInt)
	if err != nil {
		return utility.ErrorResponse("Internal Server Error", nil, ctx)
	}

	return utility.SuccessResponse("Success Deleting Contact", nil, ctx)
}
