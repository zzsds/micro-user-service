package main

import (
	"context"
	"testing"

	"github.com/micro/go-micro/v2"
	user "github.com/zzsds/micro-user-service/proto/user"
)

func TestUserClient(t *testing.T) {
	service := micro.NewService(
		micro.Name("client.user"),
	)
	service.Init()
	client := user.NewUserService("store.user", service.Client())
	result, err := client.Index(context.Background(), &user.Pagination{Size: 10})
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
