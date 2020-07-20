package service

import (
	"github.com/jinzhu/gorm"
	"github.com/zzsds/micro-store/user-service/models"
)

// UserInterface ...
type UserInterface interface {
	GetMobile(mobile string) *models.User
	GetEmail(email string) *models.User
	Create(*models.User) error
}

// User ...
type User struct {
	db *gorm.DB
}

// NewUser ...
func NewUser(DB *gorm.DB) *UserInterface {
	return &user{
		db: DB,
	}
}

// GetMobile ...
func GetMobile(mobile string) *models.User {

	return &models.User{}
}

// GetEmail ...
func GetEmail(email string) *models.User {

	return &models.User{}
}

// Create ...
func Create(*models.User) error {
	return nil
}
