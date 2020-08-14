package service

import (
	"bytes"
	"fmt"
	"time"

	"github.com/go-errors/errors"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/zzsds/micro-user-service/models"
	user "github.com/zzsds/micro-user-service/proto/user"
)

// UserInterface ...
type UserInterface interface {
	ResourceToModel(req *user.Resource) *models.User
	ModelToResource(model *models.User) *user.Resource
	PageDate(page, size int32, condition, order []string) ([]*models.User, int32)
	FindMobile(mobile string) *models.User
	FindCode(code string) *models.User
	FindEmail(email string) *models.User
	Create(*models.User) error
	FindID(id uint) *models.User
	ModifyPassword(id uint, password, old string) error
	ResetPassword(id uint, password string) error
	ModifyMobile(id uint, mobile, old string) error
	PassLogin(user, pass string) (*models.User, error)
	FindLikeMobile(mobile string) []*models.User
	FindInMobile(mobile ...string) []*models.User
	FindInID(id ...uint) []*models.User
	FindSource(source string) []*models.User
	SourceType() []string
	ModifyName(id uint, name string) error
	BatchCreate([]*models.User) error
}

// User ...
type User struct {
	*Dao
}

// NewUser ...
func NewUser(dao *Dao) UserInterface {
	return &User{
		dao,
	}
}

// ModelToResource ...
func (s *User) ModelToResource(model *models.User) *user.Resource {
	createdAt, _ := ptypes.TimestampProto(model.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(model.UpdatedAt)
	var birthday *timestamp.Timestamp
	if model.Birthday != nil {
		birthday, _ = ptypes.TimestampProto(*model.Birthday)
	}
	return &user.Resource{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Id:        int32(model.ID),
		Name:      model.Name,
		Mobile:    model.Mobile,
		Code:      model.Code,
		Source:    model.Source,
		Nickname:  model.Nickname,
		Realname:  model.Realname,
		Email:     model.Email,
		Birthday:  birthday,
		Enabled:   user.Enabled(model.Enabled),
	}
}

// ResourceToModel ...
func (s *User) ResourceToModel(req *user.Resource) *models.User {
	var birthday *time.Time
	if req.GetBirthday() != nil {
		birth, _ := ptypes.Timestamp(req.GetBirthday())
		birthday = &birth
	}
	return &models.User{
		Mobile:   req.GetMobile(),
		Name:     req.GetName(),
		Source:   req.GetSource(),
		Email:    req.GetEmail(),
		Nickname: req.GetNickname(),
		Realname: req.GetRealname(),
		Code:     req.GetCode(),
		Birthday: birthday,
		Enabled:  int32(req.GetEnabled()),
	}
}

// FindID ...
func (s *User) FindID(id uint) *models.User {
	model := models.User{}
	if s.Db().First(&model, id).RecordNotFound() {
		return nil
	}
	return &model
}

// PageDate 分页查询
func (s *User) PageDate(page, size int32, condition, order []string) (list []*models.User, total int32) {
	db := s.Db()
	for _, o := range order {
		db = db.Order(o)
	}
	for _, c := range condition {
		db = db.Where(c)
	}
	db.Model(models.User{}).Count(&total).Offset(page * size).Limit(size).Find(&list)
	return list, total
}

// FindMobile ...
func (s *User) FindMobile(mobile string) *models.User {
	model := models.User{}
	if s.Db().Where("mobile = ?", mobile).First(&model).RecordNotFound() {
		return nil
	}
	return &model
}

// FindCode ...
func (s *User) FindCode(code string) *models.User {
	model := models.User{}
	if s.Db().Where("code = ?", code).First(&model).RecordNotFound() {
		return nil
	}
	return &model
}

// FindEmail ...
func (s *User) FindEmail(email string) *models.User {
	model := models.User{}
	if s.Db().Where("email = ?", email).First(&model).RecordNotFound() {
		return nil
	}
	return &model
}

// Create ...
func (s *User) Create(model *models.User) error {
	if model.Code == "" {
		model.Code = models.GenerateCode(6)
	}

	opt := s.options
	if opt != nil {
		var i int
		if err := opt.Code.Read(model.Code, &i); err != nil {
			return s.Create(model)
		}
	} else {
		if s.FindCode(model.Code) != nil {
			return s.Create(model)
		}
	}

	tx := s.Db().Begin()

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
	if opt != nil {
		opt.Code.Write(model.Code, model.ID)
	}

	return tx.Commit().Error
}

// ModifyPassword 修改密码
func (s *User) ModifyPassword(id uint, password, oldPass string) error {
	tx := s.Db().Begin()

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

// ResetPassword 修改密码
func (s *User) ResetPassword(id uint, password string) error {
	tx := s.Db().Begin()

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

	pass, _ := models.EncodeSalt(password, model.Salt)
	if model.Password == pass {
		return fmt.Errorf("The old password is the same as the new one")
	}
	model.Password = pass
	if err := tx.Save(&model).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// ModifyMobile 修改手机号
func (s *User) ModifyMobile(id uint, mobile, oldMobile string) error {
	tx := s.Db().Begin()

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

	if model.Mobile != oldMobile {

		return fmt.Errorf("Old mobile number error")
	}

	model.Mobile = mobile
	if err := tx.Save(&model).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// ModifyName 修改名称
func (s *User) ModifyName(id uint, name string) error {
	tx := s.Db().Begin()

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

	model.Name = name
	if err := tx.Save(&model).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// PassLogin ...
func (s *User) PassLogin(name, password string) (*models.User, error) {
	model := models.User{}
	if s.Db().Where("name = ?", name).Or("mobile = ?", name).Or("email = ?", name).First(&model).RecordNotFound() {
		return nil, fmt.Errorf("user does not exist")
	}

	pass, _ := models.EncodeSalt(password, model.Salt)
	if model.Password != pass {
		return nil, fmt.Errorf("The password is wrong")
	}
	return &model, nil
}

// FindLikeMobile ...模糊查询手机号
func (s *User) FindLikeMobile(mobile string) []*models.User {
	list := make([]*models.User, 0)
	db := s.Db()
	if mobile != "" {
		db = db.Where("mobile LIKE ?", "%"+fmt.Sprintf("%s", mobile)+"%")
	}
	if db.Find(&list).RecordNotFound() {
		return nil
	}
	return list
}

// FindInMobile ...批量查询手机号
func (s *User) FindInMobile(mobile ...string) []*models.User {
	list := make([]*models.User, 0)
	db := s.Db()
	if len(mobile) > 0 {
		db = db.Where("mobile in (?)", mobile)
	}

	if db.Find(&list).RecordNotFound() {
		return nil
	}
	return list
}

// FindInID ...批量查询ID
func (s *User) FindInID(id ...uint) []*models.User {
	list := make([]*models.User, 0)
	db := s.Db()
	if len(id) > 0 {
		db = db.Where("id in (?)", id)
	}
	if db.Find(&list).RecordNotFound() {
		return nil
	}
	return list
}

// FindSource ...查询注册来源
func (s *User) FindSource(source string) []*models.User {
	list := make([]*models.User, 0)
	db := s.Db()
	if source != "" {
		db = db.Where("source = ?", source)
	}
	if db.Find(&list).RecordNotFound() {
		return nil
	}
	return list
}

// SourceType ...
func (s *User) SourceType() []string {
	model, list := models.User{}, make([]string, 0)
	if s.Db().Model(model).Select("source").Group("source").Pluck("source", &list).RecordNotFound() {
		return nil
	}
	return list
}

// BatchCreate ...
func (s *User) BatchCreate(users []*models.User) error {
	userLen := len(users)
	if userLen <= 0 {
		return errors.Errorf("user data not null")
	}
	now := time.Now()
	for _, model := range users {
		model.CreatedAt = now
		model.UpdatedAt = now
	CODE:
		if model.Code == "" {
			model.Code = models.GenerateCode(6)
		}
		model.Salt = models.EncodeMD5(model.Code)
		if model.Password != "" {
			pass, _ := models.EncodeSalt(model.Password, model.Salt)
			model.Password = pass
		}
		var i int
		s.options.Code.Read(model.Code, &i)
		if i > 0 {
			goto CODE
		}
	}

	tx := s.Db().Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	var buff bytes.Buffer
	buff.WriteString("INSERT INTO users (`created_at`, `updated_at`, `name`, `password`, `salt`, `mobile`, `email`, `nickname`, `realname`, `code`, `source`, `enabled`) VALUES ")
	localFormat := "2006-01-02 15:04:05"
	for k, v := range users {
		buff.WriteString(fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %d)", v.CreatedAt.Format(localFormat), v.UpdatedAt.Format(localFormat), v.Name, v.Password, v.Salt, v.Mobile, v.Email, v.Nickname, v.Realname, v.Code, v.Source, v.Enabled))
		divide := ","
		if k+1 == userLen {
			divide = ";"
		}
		buff.WriteString(divide)
	}
	if err := tx.Exec(buff.String()).Error; err != nil {
		tx.Rollback()
		return err
	}
	userList := make([]*models.User, 0)
	if !tx.Order("id DESC").Limit(userLen).Find(&userList).RecordNotFound() {
		copy(users, userList)
	}

	// 重新加载code列表
	go s.initCodeList(users...)

	return tx.Commit().Error
}
