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

func (repo *ContactRepository) GetAllContact(page int, size int) ([]entity.Contact, int64, error) {
	var contacts []entity.Contact
	var total int64

	// hitung total data
	err := repo.DB.Model(&entity.Contact{}).Count(&total).Error

	if err != nil {
		return nil, 0, err
	}

	// hitung offset
	offset := (page - 1) * size

	// ambil data pagination dengan limit & offset
	err = repo.DB.Limit(size).Offset(offset).Find(&contacts).Error
	if err != nil {
		return nil, 0, err
	}

	return contacts, total, nil
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

func (repo *ContactRepository) CheckExistsPhoneExceptID(phone string, id uint64) bool {
	var found bool
	err := repo.DB.Raw("SELECT EXISTS(SELECT 1 FROM contacts WHERE id <> ? AND phone = ?)", id, phone).Scan(&found).Error

	if err != nil {
		log.Printf("Error when validasi CheckExistsPhoneEdit: %v", err.Error())
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

func (repo *ContactRepository) FindContactByID(id uint64) (*entity.Contact, error) {
	var contact entity.Contact

	err := repo.DB.Model(&entity.Contact{}).Where("id = ?", id).Take(&contact).Error
	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (repo *ContactRepository) DeleteContactByID(id uint64) error {
	// Unscoped(): Menonaktifkan soft delete, sehingga record dihapus secara permanen.
	err := repo.DB.Unscoped().Where("id = ?", id).Delete(&entity.Contact{}).Error
	// err := repo.db.Exec("DELETE FROM contacts WHERE id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *ContactRepository) UpdateContactByID(id uint64, dataUpdate entity.Contact) (*entity.Contact, error) {
	var contact entity.Contact

	err := repo.DB.Model(&entity.Contact{}).Where("id = ?", id).Updates(dataUpdate).Scan(&contact).Error

	if err != nil {
		return nil, err
	}

	return &contact, nil
}
