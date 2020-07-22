package service

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/zzsds/micro-user-service/models"
)

// UserInterface ...
type UserInterface interface {
	PageDate(page, size int32, condition, order []string) ([]*models.User, int32)
	GetMobile(mobile string) *models.User
	GetEmail(email string) *models.User
	Create(*models.User) error
	FindID(id uint) *models.User
	ModifyPassword(id uint, password, old string) error
	ModifyMobile(id uint, mobile, old string) error
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
func (s *User) FindID(id uint) *models.User {
	model := models.User{}
	if s.db.First(&model, id).RecordNotFound() {
		return nil
	}
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
	model := models.User{}
	s.db.Where("mobile = ?", mobile).First(&model)
	return &model
}

// GetEmail ...
func (s *User) GetEmail(email string) *models.User {
	model := models.User{}
	s.db.Where("email = ?", email).First(&model)
	return &model
}

// Create ...
func (s *User) Create(model *models.User) error {
	tx := s.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(model).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// ModifyPassword 修改密码
func (s *User) ModifyPassword(id uint, password, oldPass string) error {
	tx := s.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	model := s.FindID(id)
	if model == nil {
		return fmt.Errorf("User Query Failed")
	}

	if model.Password == "" {
		return fmt.Errorf("Password not set")
	}

	old, _ := models.EncodeSalt(oldPass, model.Salt)
	if model.Password != old {
		return fmt.Errorf("Old password validation failed")
	}
	pass, _ := models.EncodeSalt(password, model.Salt)
	model.Password = pass
	if err := tx.Save(&model).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// ModifyMobile 修改手机号
func (s *User) ModifyMobile(id uint, mobile, oldPass string) error {
	tx := s.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	model := s.FindID(id)
	if model == nil {
		return fmt.Errorf("User Query Failed")
	}

	if model.Mobile == "" {
		return fmt.Errorf("Mobile not set")
	}

	if model.Mobile != oldPass {

		return fmt.Errorf("Old mobile number error")
	}

	model.Mobile = mobile
	if err := tx.Save(&model).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
