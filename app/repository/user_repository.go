package repository

import (
	"cobaaja/contact-management/app/dto"
	"cobaaja/contact-management/app/entity"
	"cobaaja/contact-management/utility"
	"log"
	"time"

	"gorm.io/gorm"
)

// Repository Implementation
type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repo *UserRepository) FindByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := repo.DB.Model(entity.User{}).Where("username = ?", username).Take(&user).Error
	return &user, err
}

func (repo *UserRepository) CheckExistUsername(username string) bool {
	var found bool
	err := repo.DB.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", username).Scan(&found).Error

	if err != nil {
		log.Printf("Error when validasi CheckExistUsername: %v", err.Error())
		return true
	}

	return found
}

func (repo *UserRepository) CreateUser(tx *gorm.DB, dto *dto.RegisterRequest) (*entity.User, error) {
	var err error
	bcryptPassword, err := utility.HashPassword(dto.Password)

	if err != nil {
		return nil, err
	}

	// konversi dari dto ke model
	newUser := &entity.User{
		Username:  dto.Username,
		Password:  bcryptPassword,
		Name:      dto.Name,
		CreatedAt: time.Now(),
	}

	err = tx.Create(&newUser).Error

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return newUser, nil
}
