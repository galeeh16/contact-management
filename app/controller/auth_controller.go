package controller

import (
	"cobaaja/contact-management/app/dto"
	"cobaaja/contact-management/app/repository"
	"cobaaja/contact-management/utility"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	Repo *repository.UserRepository
}

func NewAuthController(repo *repository.UserRepository) *AuthController {
	return &AuthController{Repo: repo}
}

func (ctrl *AuthController) Register(ctx *fiber.Ctx) error {
	req := new(dto.RegisterRequest)

	// binding request ke struct
	ctx.BodyParser(&req)

	// validasi requestnya
	v := utility.NewValidator()

	// validasi unique username
	v.Validate.RegisterValidation("unique_username", func(fl validator.FieldLevel) bool {
		existsUsername := ctrl.Repo.CheckExistUsername(req.Username)
		return !existsUsername
	})

	// validasi strong password
	v.Validate.RegisterValidation("strong_password", utility.RegisterStrongPasswordValidation)

	arrayError := v.ValidateStruct(req)

	if arrayError != nil {
		return utility.BadRequestResponse("Invalid Data", arrayError, ctx)
	}

	// begin transaction
	tx := ctrl.Repo.DB.Begin()

	// create user
	newUser, err := ctrl.Repo.CreateUser(tx, req)
	if err != nil {
		return utility.ErrorResponse("Failed to Create User", err.Error(), ctx)
	}

	// mapping dari model ke dto response
	userRes := &dto.RegisterResponse{
		ID:        newUser.ID,
		Username:  newUser.Username,
		Name:      newUser.Name,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}

	// commit kalo sukses
	tx.Commit()

	return utility.SuccessResponse("Register Success", userRes, ctx)
}

func (ctrl *AuthController) Login(ctx *fiber.Ctx) error {
	req := new(dto.LoginRequest)

	// bind request ke struct
	ctx.BodyParser(&req)

	// validasi login request
	v := utility.NewValidator()

	arrayError := v.ValidateStruct(req)

	if arrayError != nil {
		return utility.BadRequestResponse("Invalid Data", arrayError, ctx)
	}

	// cek username exists
	user, err := ctrl.Repo.FindByUsername(req.Username)

	// kalo error, karena record tidak ketemu atau error kodingan
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utility.BadRequestResponse("Username or password are incorrect", nil, ctx)
		} else {
			return utility.ErrorResponse("Internal Server Error", err.Error(), ctx)
		}
	}

	// check passwordnya bener ga
	err = utility.VerifyPassword(user.Password, req.Password)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return utility.BadRequestResponse("Username or password are incorrect", nil, ctx)
		}

		return utility.ErrorResponse("Internal Server Error", err.Error(), ctx)
	}

	// create token jwt
	jwtToken, err := utility.CreateJwtToken(user)
	if err != nil {
		return utility.ErrorResponse("Failed to create token", nil, ctx)
	}

	return utility.SuccessResponse("Login Success", fiber.Map{
		"access_token": jwtToken,
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"name":     user.Name,
		},
	}, ctx)
}

func (ctrl *AuthController) Me(ctx *fiber.Ctx) error {
	// ambil username dari ctx local yang dibuat di jwt middleware
	username, ok := ctx.Locals("username").(string)
	if !ok {
		return utility.BadRequestResponse("Username tidak valid", nil, ctx)
	}

	// get user by username
	user, err := ctrl.Repo.FindByUsername(username)
	if err != nil {
		return utility.BadRequestResponse("Username "+username+" not found", nil, ctx)
	}

	// mapping user ke dto
	response := &dto.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return utility.SuccessResponse("Success", response, ctx)
}
