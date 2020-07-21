package handler

import (
	"context"

	"github.com/micro/go-micro/v2"
	user "github.com/zzsds/micro-store/user-service/proto/user"
	"github.com/zzsds/micro-store/user-service/service"
)

// User ...
type User struct {
	name string
}

// NewUserHandler ...初始化Handler
func NewUserHandler(srv micro.Service, db *service.Dao) *User {
	return &User{
		name: srv.Name(),
	}
}

// Index ...
func (h *User) Index(ctx context.Context, req *user.Pagination, rsp *user.List) error {
	return nil
}

// Show ...
func (h *User) Show(ctx context.Context, req *user.ShowRequest, rsp *user.ShowResponse) error {
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
