package repository

import (
	"cobaaja/contact-management/app/dto"
	"cobaaja/contact-management/app/entity"
	"cobaaja/contact-management/utility"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Repository Implementation
type UserRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewUserRepository(db *gorm.DB, logger *logrus.Logger) *UserRepository {
	return &UserRepository{
		DB:     db,
		Logger: logger,
	}
}

func (repo *UserRepository) CreateJwtToken(user *entity.User) (string, error) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	var ttlMinute, _ = strconv.ParseInt(os.Getenv("JWT_TTL"), 10, 64)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"user_id":  user.ID,
		"exp":      time.Now().Add(time.Hour * time.Duration(ttlMinute)).Unix(), // Token berlaku 2 jam
	})
	return token.SignedString(jwtSecret)
}

func (repo *UserRepository) CreateJwtRefreshToken(user *entity.User) (string, error) {
	var jwtRefreshSecret = []byte(os.Getenv("JWT_REFRESH_SECRET"))
	var ttlMinute, _ = strconv.ParseInt(os.Getenv("JWT_REFRESH_TTL"), 10, 64)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   user.Username,
		"user_id":    user.ID,
		"is_refresh": true,
		"exp":        time.Now().Add(time.Minute * time.Duration(ttlMinute)).Unix(), // Token berlaku 1 minggu
	})
	return token.SignedString(jwtRefreshSecret)
}

func (repo *UserRepository) FindByUsername(username string) (*entity.User, error) {
	repo.Logger.Info("Find by username: " + username)
	var user entity.User
	err := repo.DB.Model(entity.User{}).Where("username = ?", username).Take(&user).Error
	return &user, err
}

func (repo *UserRepository) CheckExistUsername(username string) bool {
	repo.Logger.Info("Check exist by username: " + username)
	var found bool
	err := repo.DB.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", username).Scan(&found).Error

	if err != nil {
		log.Printf("Error when validasi CheckExistUsername: %v", err.Error())
		return true
	}

	return found
}

func (repo *UserRepository) CreateUser(tx *gorm.DB, dto *dto.RegisterRequest) (*entity.User, error) {
	repo.Logger.Info("Create User")
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
