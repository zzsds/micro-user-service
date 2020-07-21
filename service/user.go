package service

import (
	"github.com/jinzhu/gorm"
	"github.com/zzsds/micro-store/user-service/models"
)

// UserInterface ...
type UserInterface interface {
	PageDate(page, size int32, condition, order []string) ([]*models.User, int32)
	GetMobile(mobile string) *models.User
	GetEmail(email string) *models.User
	Create(*models.User) error
	FindID(id int32) *models.User
}

// User ...
type User struct {
	db *gorm.DB
}

// NewUser ...
func NewUser(DB *gorm.DB) UserInterface {
	return &User{
		db: DB,
	}
}

// FindID ...
func (s *User) FindID(id int32) *models.User {
	model := models.User{}
	s.db.First(&model)
	return &model
}

// PageDate 分页查询
func (s *User) PageDate(page, size int32, condition, order []string) (list []*models.User, total int32) {
	for _, o := range order {
		s.db = s.db.Order(o)
	}
	for _, c := range condition {
		s.db = s.db.Where(c)
	}
	s.db.Model(models.User{}).Count(&total).Offset(page * size).Limit(size).Find(&list)
	return list, total
}

// GetMobile ...
func (s *User) GetMobile(mobile string) *models.User {

	return &models.User{}
}

// GetEmail ...
func (s *User) GetEmail(email string) *models.User {

	return &models.User{}
}

// Create ...
func (s *User) Create(*models.User) error {
	return nil
}
