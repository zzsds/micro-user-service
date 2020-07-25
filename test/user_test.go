package main

import (
	"context"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/micro/go-micro/v2"
	user "github.com/zzsds/micro-user-service/proto/user"
)

var service micro.Service

func TestUserIndex(t *testing.T) {
	service = micro.NewService(
		micro.Name("client.user"),
	)

	client := user.NewUserService("store.srv.user", service.Client())
	result, err := client.Index(context.Background(), &user.Pagination{Size: 10})
	if err != nil {
		t.Error(err)
	}
	t.Log(result.GetData())
}

func TestUserSearchPage(t *testing.T) {
	service = micro.NewService(
		micro.Name("client.user"),
	)
	client := user.NewUserService("store.srv.user", service.Client())
	// 一个月前未开始时间
	start, _ := ptypes.TimestampProto(time.Now().AddDate(0, -1, 0))
	end := ptypes.TimestampNow()
	t.Log(start, end)
	var createAt *user.Between
	if start != nil && end != nil {
		createAt = &user.Between{
			Start: start,
			End:   ptypes.TimestampNow(),
		}
	}
	t.Log(createAt)
	result, err := client.SearchPage(context.Background(), &user.SearchPageRequest{
		Condition: &user.SearchPageRequest_Condition{
			Mobile:    "1",
			CreatedAt: createAt,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.Total)
}
