package repository

import (
	"cobaaja/contact-management/app/entity"
	"log"

	"gorm.io/gorm"
)

type ContactRepository struct {
	DB *gorm.DB
}

func NewContactRepository(db *gorm.DB) *ContactRepository {
	return &ContactRepository{DB: db}
}

func (repo *ContactRepository) CreateContact(tx *gorm.DB, newContact *entity.Contact) (*entity.Contact, error) {
	var contact entity.Contact

	err := tx.Create(&newContact).Scan(&contact).Error
	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (repo *ContactRepository) CheckExistsPhone(phone string) bool {
	var found bool
	err := repo.DB.Raw("SELECT EXISTS(SELECT 1 FROM contacts WHERE phone = ?)", phone).Scan(&found).Error

	if err != nil {
		log.Printf("Error when validasi CheckExistsPhone: %v", err.Error())
		return true
	}

	return found
}

func (repo *ContactRepository) FindContactByPhone(phone string) (*entity.Contact, error) {
	var contact entity.Contact

	err := repo.DB.Model(&entity.Contact{}).Where("phone = ?", phone).Take(&contact).Error
	if err != nil {
		return nil, err
	}

	return &contact, nil
}
