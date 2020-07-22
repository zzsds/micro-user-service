package handler

import (
	"context"
	"strings"

	"github.com/golang/protobuf/ptypes"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/errors"
	"github.com/zzsds/micro-user-service/models"
	user "github.com/zzsds/micro-user-service/proto/user"
	"github.com/zzsds/micro-user-service/service"
)

const (
	passLen = 6
)

// User ...
type User struct {
	dao     *service.Dao
	name    string
	service service.UserInterface
}

// NewUserHandler ...初始化Handler
func NewUserHandler(srv micro.Service, dao *service.Dao) *User {
	return &User{
		dao:     dao,
		name:    srv.Name(),
		service: service.NewUser(dao.Db()),
	}
}

func (h *User) String(params ...string) string {
	return h.name + " User." + strings.Join(params, " ")
}

// Index ...
func (h *User) Index(ctx context.Context, req *user.Pagination, rsp *user.List) error {
	if req.GetSize() <= 0 {
		req.Size = 20
	}
	list, total := h.service.PageDate(req.GetPage(), req.GetSize(), req.GetCondition(), req.GetOrder())
	req.Total = total

	for _, model := range list {
		createdAt, _ := ptypes.TimestampProto(model.CreatedAt)
		updatedAt, _ := ptypes.TimestampProto(model.UpdatedAt)
		resource := &user.Resource{
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Id:        int32(model.ID),
			Name:      model.Name,
			Mobile:    model.Mobile,
			Code:      model.Code,
			Source:    model.Source,
			Enabled:   user.Enabled(model.Enabled),
		}
		if model.Birthday != nil {
			resource.Birthday, _ = ptypes.TimestampProto(*model.Birthday)
		}
		rsp.Data = append(rsp.GetData(), resource)
	}

	return nil
}

// Show ...
func (h *User) Show(ctx context.Context, req *user.ShowRequest, rsp *user.ShowResponse) error {
	if req.GetId() <= 0 {
		return errors.BadRequest(h.String("Show"), "ID 不能为空")
	}
	model := h.service.FindID(uint(req.GetId()))
	if model == nil {
		return nil
	}
	createdAt, _ := ptypes.TimestampProto(model.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(model.UpdatedAt)
	rsp.Data = &user.Resource{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Id:        int32(model.ID),
		Name:      model.Name,
		Mobile:    model.Mobile,
		Code:      model.Code,
		Source:    model.Source,
		Enabled:   user.Enabled(model.Enabled),
	}
	if model.Birthday != nil {
		rsp.Data.Birthday, _ = ptypes.TimestampProto(*model.Birthday)
	}
	return nil
}

// GetMobile ...
func (h *User) GetMobile(ctx context.Context, req *user.MobileRequest, rsp *user.MobileResponse) error {
	if !models.ValidateMobile(req.GetMobile()) {
		return errors.BadRequest(h.String("GetMobile"), "Mobile 格式错误")
	}
	model := h.service.GetMobile(req.GetMobile())
	if model == nil {
		return nil
	}
	createdAt, _ := ptypes.TimestampProto(model.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(model.UpdatedAt)

	rsp.Data = &user.Resource{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Id:        int32(model.ID),
		Name:      model.Name,
		Mobile:    model.Mobile,
		Code:      model.Code,
		Source:    model.Source,
		Enabled:   user.Enabled(model.Enabled),
	}

	if model.Birthday != nil {
		rsp.Data.Birthday, _ = ptypes.TimestampProto(*model.Birthday)
	}
	return nil
}

// MobileRegister ...
func (h *User) MobileRegister(ctx context.Context, req *user.MobileRegisterRequest, rsp *user.MobileRegisterResponse) error {
	model := models.User{
		Mobile:   req.GetMobile(),
		Password: req.GetPassword(),
		Name:     req.GetName(),
		Source:   req.GetSource(),
	}
	if model.Source == "" {
		return errors.BadRequest(h.String("MobileRegister"), "创建来源不能为空")
	}
	if !models.ValidateMobile(model.Mobile) {
		return errors.BadRequest(h.String("MobileRegister"), "手机号格式错误")
	}
	if model.Name == "" {
		model.Name = model.Mobile
	}

	if h.service.GetMobile(model.Mobile).ID > 0 {
		return errors.BadRequest(h.String("MobileRegister"), "%s 手机号已存在", model.Mobile)
	}

	if len(req.GetPassword()) < passLen {
		return errors.BadRequest(h.String("MobileRegister"), "密码长度不能小于%d", passLen)
	}

	err := h.service.Create(&model)
	if err != nil {
		return errors.BadRequest(h.String("MobileRegister"), "数据保存失败：%s", err.Error())
	}
	rsp.Id = int32(model.ID)
	return nil
}

// ModifyPassword ...
func (h *User) ModifyPassword(ctx context.Context, req *user.ModifyPassRequest, rsp *user.ModifyPassResponse) error {
	if req.GetId() <= 0 {
		return errors.BadRequest(h.String("ModifyPassword"), "UID 不能为空")
	}

	if req.GetPassword() == req.GetOldPassword() {
		return errors.BadRequest(h.String("ModifyPassword"), "新密码和旧密码一致")
	}

	if len(req.GetPassword()) < passLen {
		return errors.BadRequest(h.String("ModifyPassword"), "密码长度不能小于%d", passLen)
	}

	if err := h.service.ModifyPassword(uint(req.Id), req.GetPassword(), req.GetOldPassword()); err != nil {
		return errors.BadRequest(h.String("ModifyPassword"), "修改失败：%v", err)
	}

	rsp.Success = true
	return nil
}

// ResetPassword ...
func (h *User) ResetPassword(ctx context.Context, req *user.ResetPassRequest, rsp *user.ResetPassResponse) error {
	if req.GetId() <= 0 {
		return errors.BadRequest(h.String("ModifyPassword"), "UID 不能为空")
	}

	if len(req.GetPassword()) < passLen {
		return errors.BadRequest(h.String("ModifyPassword"), "密码长度不能小于%d", passLen)
	}

	if err := h.service.ResetPassword(uint(req.Id), req.GetPassword()); err != nil {
		return errors.BadRequest(h.String("ModifyPassword"), "重置失败：%v", err)
	}

	rsp.Success = true
	return nil
}

// ModifyMobile ...
func (h *User) ModifyMobile(ctx context.Context, req *user.ModifyMobileRequest, rsp *user.ModifyMobileResponse) error {
	if req.GetId() <= 0 {
		return errors.BadRequest(h.String("ModifyMobile"), "UID 不能为空")
	}
	if !models.ValidateMobile(req.GetMobile()) {
		return errors.BadRequest(h.String("ModifyPassword"), "手机号格式错误")
	}
	if !models.ValidateMobile(req.GetOldMobile()) {
		return errors.BadRequest(h.String("ModifyPassword"), "旧手机号格式错误")
	}
	if err := h.service.ModifyMobile(uint(req.GetId()), req.GetMobile(), req.GetOldMobile()); err != nil {
		return errors.BadRequest(h.String("ModifyMobile"), "修改失败：%v", err)
	}
	rsp.Success = true
	return nil
}

// PassLogin ...
func (h *User) PassLogin(ctx context.Context, req *user.PassLoginRequest, rsp *user.PassLoginResponse) error {
	if req.GetUser() == "" {
		return errors.BadRequest(h.String("PassLogin"), "账号不能为空")
	}

	if len(req.GetPassword()) < passLen {
		return errors.BadRequest(h.String("PassLogin"), "密码长度不能小于%d", passLen)
	}

	user, err := h.service.PassLogin(req.GetUser(), req.GetPassword())
	if err != nil {
		return errors.BadRequest(h.String("PassLogin"), "登录失败：%v", err)
	}

	rsp.Id = int32(user.ID)
	rsp.Success = true

	return nil
}
