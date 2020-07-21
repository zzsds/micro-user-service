package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/micro/go-micro/v2"
	user "github.com/zzsds/micro-store/user-service/proto/user"
	"github.com/zzsds/micro-store/user-service/service"
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
		birthday, _ := ptypes.TimestampProto(*model.Birthday)
		rsp.Data = append(rsp.GetData(), &user.Resource{
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Id:        int32(model.ID),
			Name:      model.Name,
			Mobile:    model.Mobile,
			Code:      model.Code,
			Source:    model.Source,
			Birthday:  birthday,
			Enabled:   user.Enabled(model.Enabled),
		})
	}

	return nil
}

// Show ...
func (h *User) Show(ctx context.Context, req *user.ShowRequest, rsp *user.ShowResponse) error {
	if req.GetId() <= 0 {
		// return errors.BadRequest()
	}
	model := h.service.FindID(req.GetId())
	if model == nil {
		return nil
	}
	createdAt, _ := ptypes.TimestampProto(model.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(model.UpdatedAt)
	birthday, _ := ptypes.TimestampProto(*model.Birthday)
	*rsp.Data = user.Resource{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Id:        int32(model.ID),
		Name:      model.Name,
		Mobile:    model.Mobile,
		Code:      model.Code,
		Source:    model.Source,
		Birthday:  birthday,
		Enabled:   user.Enabled(model.Enabled),
	}
	return nil
}

// GetMobile ...
func (h *User) GetMobile(ctx context.Context, req *user.MobileRequest, rsp *user.MobileResponse) error {
	return nil
}

// MobileCreate ...
func (h *User) MobileCreate(ctx context.Context, req *user.MobileCreateRequest, rsp *user.MobileCreateResponse) error {
	return nil
}

// ModifyPassword ...
func (h *User) ModifyPassword(ctx context.Context, req *user.ModifyPassRequest, rsp *user.ModifyPassResponse) error {
	return nil
}

// ModifyMobile ...
func (h *User) ModifyMobile(ctx context.Context, req *user.ModifyPassRequest, rsp *user.ModifyPassResponse) error {
	return nil
}
