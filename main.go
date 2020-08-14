package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config/source/file"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/store/memory"
	"github.com/zzsds/micro-user-service/conf"
	"github.com/zzsds/micro-user-service/handler"
	dao "github.com/zzsds/micro-user-service/service"
	"github.com/zzsds/micro-user-service/subscriber"
	"github.com/zzsds/micro-utils/config/nacos"

	user "github.com/zzsds/micro-user-service/proto/user"
)

const (
	name    = "store.srv.user"
	version = "latest"
)

func main() {
	conf.InitConfig(file.WithPath("config.toml"), nacos.WithDataIDKey(name))
	// New Service
	service := micro.NewService(
		micro.Name(name),
		micro.Version(version),
	)

	// Initialise service
	service.Init()
	dao := dao.NewDao(dao.WithStore(memory.NewStore()))
	defer dao.Close()

	// Register Handler
	user.RegisterUserHandler(service.Server(), handler.NewUserHandler(service, dao))

	// Register Struct as Subscriber
	micro.RegisterSubscriber(name, service.Server(), new(subscriber.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
