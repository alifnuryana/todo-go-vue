package repository

import (
	"github.com/alifnuryana/go-auth-jwt/models"
	"gorm.io/gorm"
)

type UserInterface interface {
	GetUserByEmail(email string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	CreateUser(user models.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserInterface {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	tx := u.db.Where("username = ?", username).First(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func (u *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	tx := u.db.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func (u *UserRepository) CreateUser(user models.User) error {
	tx := u.db.Create(&user)
	return tx.Error
}
