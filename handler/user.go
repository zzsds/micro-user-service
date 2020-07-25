package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-playground/locales/zh"
	"github.com/golang/protobuf/ptypes"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/errors"
	"github.com/zzsds/micro-user-service/models"
	user "github.com/zzsds/micro-user-service/proto/user"
	"github.com/zzsds/micro-user-service/service"
)

// UserHandler ...
type UserHandler interface {
	Index(context.Context, *user.Pagination, *user.List) error
	Show(context.Context, *user.ShowRequest, *user.ShowResponse) error
	GetMobile(context.Context, *user.MobileRequest, *user.MobileResponse) error
	MobileRegister(context.Context, *user.MobileRegisterRequest, *user.MobileRegisterResponse) error
	ModifyPassword(context.Context, *user.ModifyPassRequest, *user.ModifyPassResponse) error
	ResetPassword(context.Context, *user.ResetPassRequest, *user.ResetPassResponse) error
	ModifyMobile(context.Context, *user.ModifyMobileRequest, *user.ModifyMobileResponse) error
	PassLogin(context.Context, *user.PassLoginRequest, *user.PassLoginResponse) error
	FindMobile(context.Context, *user.FindMobileRequest, *user.MobileResponse) error
	FindLikeMobileList(context.Context, *user.FindLikeMobileRequest, *user.List) error
	FindInMobileList(context.Context, *user.FindInMobileRequest, *user.List) error
	FindInIDList(context.Context, *user.FindInIdRequest, *user.List) error
	FindSourceList(context.Context, *user.FindSourceRequest, *user.List) error
	SourceTypeList(context.Context, *user.SourceTypeRequest, *user.SourceTypeResponse) error
	SearchPage(ctx context.Context, in *user.SearchPageRequest, out *user.List) error
}

const (
	passLen = 6
)

// User ...
type User struct {
	validate *service.Valid
	name     string
	service  service.UserInterface
}

// NewUserHandler ...初始化Handler
func NewUserHandler(srv micro.Service, dao *service.Dao) UserHandler {
	return &User{
		name:     srv.Name(),
		service:  service.NewUser(dao),
		validate: service.NewValid(zh.New()),
	}
}

func (h *User) String(params ...string) string {
	return h.name + " User." + strings.Join(params, " ")
}

// SearchPage ...
func (h *User) SearchPage(ctx context.Context, req *user.SearchPageRequest, rsp *user.List) error {
	if req.GetSize() <= 0 {
		req.Size = 20
	}

	condition := make([]string, 0)
	if cd := req.GetCondition(); cd != nil {
		if cd.GetMobile() != "" {
			condition = append(condition, fmt.Sprintf("mobile LIKE '%s'", "%"+cd.Mobile+"%"))
		}
		if create := cd.GetCreatedAt(); create != nil {
			start := ptypes.TimestampString(create.GetStart())
			end := ptypes.TimestampString(create.GetStart())
			condition = append(condition, fmt.Sprintf("created_at BETWEEN '%s' AND '%s'", start, end))
		}
	}
	order := make([]string, len(req.GetOrder()))
	for k, v := range req.GetOrder() {
		order[k] = fmt.Sprintf("%s %s", v.GetKey(), v.GetVal().String())
	}

	list, total := h.service.PageDate(req.GetPage(), req.GetSize(), condition, order)
	rsp.Total = total
	for _, model := range list {
		rsp.Data = append(rsp.GetData(), h.service.ModelToResource(model))
	}

	return nil
}

// Index ...
func (h *User) Index(ctx context.Context, req *user.Pagination, rsp *user.List) error {
	if req.GetSize() <= 0 {
		req.Size = 20
	}
	list, total := h.service.PageDate(req.GetPage(), req.GetSize(), req.GetCondition(), req.GetOrder())
	rsp.Total = total

	for _, model := range list {
		rsp.Data = append(rsp.GetData(), h.service.ModelToResource(model))
	}

	return nil
}

// Show ...
func (h *User) Show(ctx context.Context, req *user.ShowRequest, rsp *user.ShowResponse) error {
	if err := h.validate.FirstError(h.validate.Var(req.GetId(), "required,gte=0")); err != nil {
		return errors.BadRequest(h.String("Show"), fmt.Sprintf("ID %s", err.Error()))
	}
	model := h.service.FindID(uint(req.GetId()))
	if model == nil {
		return nil
	}
	rsp.Data = h.service.ModelToResource(model)
	return nil
}

// GetMobile ...
func (h *User) GetMobile(ctx context.Context, req *user.MobileRequest, rsp *user.MobileResponse) error {
	if !models.ValidateMobile(req.GetMobile()) {
		return errors.BadRequest(h.String("GetMobile"), "%s 手机号格式错误", req.GetMobile())
	}
	model := h.service.FindMobile(req.GetMobile())
	if model == nil {
		return nil
	}

	rsp.Data = h.service.ModelToResource(model)
	return nil
}

// MobileRegister ...
func (h *User) MobileRegister(ctx context.Context, req *user.MobileRegisterRequest, rsp *user.MobileRegisterResponse) error {
	if !models.ValidateMobile(req.GetMobile()) {
		return errors.BadRequest(h.String("MobileRegister"), "手机号格式错误")
	}

	if h.service.FindMobile(req.GetMobile()).ID > 0 {
		return errors.BadRequest(h.String("MobileRegister"), "%s 手机号已存在", req.GetMobile())
	}

	if err := h.validate.FirstError(h.validate.Var(req.GetPassword(), fmt.Sprintf("required,gte=%d", passLen))); err != nil {
		return errors.BadRequest(h.String("MobileRegister"), fmt.Sprintf("Password %s", err.Error()))
	}

	if err := h.validate.FirstError(h.validate.Var(req.GetSource(), "required")); err != nil {
		return errors.BadRequest(h.String("MobileRegister"), fmt.Sprintf("Source %s", err.Error()))
	}

	if req.GetName() == "" {
		req.Name = req.GetMobile()
	}

	if err := h.validate.FirstError(h.validate.Var(req.GetName(), "required")); err != nil {
		return errors.BadRequest(h.String("MobileRegister"), err.Error())
	}

	model := models.User{
		Mobile:   req.GetMobile(),
		Password: req.GetPassword(),
		Name:     req.GetName(),
		Source:   req.GetSource(),
	}

	if err := h.service.Create(&model); err != nil {
		return errors.BadRequest(h.String("MobileRegister"), "数据保存失败：%s", err.Error())
	}
	rsp.Id = int32(model.ID)
	rsp.Success = true
	return nil
}

// ModifyPassword ...
func (h *User) ModifyPassword(ctx context.Context, req *user.ModifyPassRequest, rsp *user.ModifyPassResponse) error {

	if err := h.validate.FirstError(h.validate.Var(req.GetId(), "required,gte=0")); err != nil {
		return errors.BadRequest(h.String("ModifyPassword"), fmt.Sprintf("ID %s", err.Error()))
	}

	if req.GetPassword() == req.GetOldPassword() {
		return errors.BadRequest(h.String("ModifyPassword"), "新密码和旧密码一致")
	}

	if err := h.validate.FirstError(h.validate.Var(req.GetOldPassword(), fmt.Sprintf("required,gte=%d", passLen))); err != nil {
		return errors.BadRequest(h.String("ModifyPassword"), fmt.Sprintf("OldPassword %s", err.Error()))
	}

	if err := h.validate.FirstError(h.validate.Var(req.GetPassword(), fmt.Sprintf("required,gte=%d", passLen))); err != nil {
		return errors.BadRequest(h.String("ModifyPassword"), fmt.Sprintf("Password %s", err.Error()))
	}

	if err := h.service.ModifyPassword(uint(req.Id), req.GetPassword(), req.GetOldPassword()); err != nil {
		return errors.BadRequest(h.String("ModifyPassword"), "修改失败：%v", err)
	}

	rsp.Success = true
	return nil
}

// ResetPassword ...
func (h *User) ResetPassword(ctx context.Context, req *user.ResetPassRequest, rsp *user.ResetPassResponse) error {
	if err := h.validate.FirstError(h.validate.Var(req.GetId(), "required,gte=0")); err != nil {
		return errors.BadRequest(h.String("ResetPassword"), fmt.Sprintf("ID %s", err.Error()))
	}

	if err := h.validate.FirstError(h.validate.Var(req.GetPassword(), fmt.Sprintf("required,gte=%d", passLen))); err != nil {
		return errors.BadRequest(h.String("ResetPassword"), fmt.Sprintf("Password %s", err.Error()))
	}

	if err := h.service.ResetPassword(uint(req.Id), req.GetPassword()); err != nil {
		return errors.BadRequest(h.String("ResetPassword"), "重置失败：%v", err)
	}

	rsp.Success = true
	return nil
}

// ModifyMobile ...
func (h *User) ModifyMobile(ctx context.Context, req *user.ModifyMobileRequest, rsp *user.ModifyMobileResponse) error {
	if err := h.validate.FirstError(h.validate.Var(req.GetId(), "required,gte=0")); err != nil {
		return errors.BadRequest(h.String("ModifyMobile"), fmt.Sprintf("ID %s", err.Error()))
	}
	if !models.ValidateMobile(req.GetMobile()) {
		return errors.BadRequest(h.String("ModifyMobile"), "手机号格式错误")
	}
	if !models.ValidateMobile(req.GetOldMobile()) {
		return errors.BadRequest(h.String("ModifyMobile"), "旧手机号格式错误")
	}
	if err := h.service.ModifyMobile(uint(req.GetId()), req.GetMobile(), req.GetOldMobile()); err != nil {
		return errors.BadRequest(h.String("ModifyMobile"), "修改失败：%v", err)
	}
	rsp.Success = true
	return nil
}

// PassLogin ...
func (h *User) PassLogin(ctx context.Context, req *user.PassLoginRequest, rsp *user.PassLoginResponse) error {
	if err := h.validate.FirstError(h.validate.Var(req.GetUser(), "required")); err != nil {
		return errors.BadRequest(h.String("PassLogin"), fmt.Sprintf("User %s", err.Error()))
	}

	if err := h.validate.FirstError(h.validate.Var(req.GetPassword(), fmt.Sprintf("required,gte=%d", passLen))); err != nil {
		return errors.BadRequest(h.String("PassLogin"), fmt.Sprintf("Password %s", err.Error()))
	}

	user, err := h.service.PassLogin(req.GetUser(), req.GetPassword())
	if err != nil {
		return errors.BadRequest(h.String("PassLogin"), "登录失败：%v", err)
	}

	rsp.Id = int32(user.ID)
	rsp.Success = true

	return nil
}

// FindMobile ...
func (h *User) FindMobile(ctx context.Context, req *user.FindMobileRequest, rsp *user.MobileResponse) error {
	if !models.ValidateMobile(req.GetMobile()) {
		return errors.BadRequest(h.String("GetMobile"), "%s 手机号格式错误", req.GetMobile())
	}
	model := h.service.FindMobile(req.GetMobile())
	if model == nil {
		return nil
	}

	rsp.Data = h.service.ModelToResource(model)
	return nil
}

// FindLikeMobileList 模糊查询手机号
func (h *User) FindLikeMobileList(ctx context.Context, req *user.FindLikeMobileRequest, rsp *user.List) error {
	if err := h.validate.FirstError(h.validate.Var(req.GetMobile(), "required,alphanum")); err != nil {
		return errors.BadRequest(h.String("FindLikeMobileList"), fmt.Sprintf("Mobile %s", err.Error()))
	}
	list := h.service.FindLikeMobile(req.GetMobile())

	for _, model := range list {
		rsp.Data = append(rsp.GetData(), h.service.ModelToResource(model))
	}
	return nil
}

// FindInMobileList ...
func (h *User) FindInMobileList(ctx context.Context, req *user.FindInMobileRequest, rsp *user.List) error {
	if err := h.validate.FirstError(h.validate.Var(req.GetMobile(), "required,dive,required,alphanum")); err != nil {
		return errors.BadRequest(h.String("FindInMobileList"), fmt.Sprintf("Mobile %s", err.Error()))
	}

	mapMbs, mbs := make(map[string]string, 0), []string{}
	for _, mobile := range req.GetMobile() {
		if _, ok := mapMbs[mobile]; ok {
			continue
		}
		mapMbs[mobile] = mobile
		mbs = append(mbs, mobile)
	}

	list := h.service.FindInMobile(mbs...)

	for _, model := range list {
		rsp.Data = append(rsp.GetData(), h.service.ModelToResource(model))
	}

	return nil
}

// FindInIDList ...
func (h *User) FindInIDList(ctx context.Context, req *user.FindInIdRequest, rsp *user.List) error {
	if err := h.validate.FirstError(h.validate.Var(req.GetId(), "required,dive,required")); err != nil {
		return errors.BadRequest(h.String("FindInIDList"), fmt.Sprintf("ID %s", err.Error()))
	}

	mapIds, ids := make(map[uint]uint, 0), []uint{}
	for _, id := range req.GetId() {
		if _, ok := mapIds[uint(id)]; ok {
			continue
		}
		mapIds[uint(id)] = uint(id)
		ids = append(ids, uint(id))
	}

	list := h.service.FindInID(ids...)

	for _, model := range list {
		rsp.Data = append(rsp.GetData(), h.service.ModelToResource(model))
	}
	return nil
}

// FindSourceList ...
func (h *User) FindSourceList(ctx context.Context, req *user.FindSourceRequest, rsp *user.List) error {
	if err := h.validate.FirstError(h.validate.Var(req.GetSource(), "required")); err != nil {
		return errors.BadRequest(h.String("FindSourceList"), fmt.Sprintf("Source %s", err.Error()))
	}

	list := h.service.FindSource(req.GetSource())
	for _, model := range list {
		rsp.Data = append(rsp.GetData(), h.service.ModelToResource(model))
	}

	return nil
}

// SourceTypeList ...
func (h *User) SourceTypeList(ctx context.Context, req *user.SourceTypeRequest, rsp *user.SourceTypeResponse) error {
	rsp.Type = h.service.SourceType()
	return nil
}
