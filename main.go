package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config/source/file"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/zzsds/micro-user-service/conf"
	"github.com/zzsds/micro-user-service/handler"
	"github.com/zzsds/micro-user-service/service"
	"github.com/zzsds/micro-user-service/subscriber"
	"github.com/zzsds/micro-utils/config/nacos"

	user "github.com/zzsds/micro-user-service/proto/user"
)

const (
	name    = "store.user"
	version = "latest"
)

func main() {
	conf.InitConfig(file.WithPath("config.toml"), nacos.WithDataIDKey(name))
	// New Service
	srv := micro.NewService(
		micro.Name(name),
		micro.Version(version),
	)

	// Initialise service
	srv.Init()
	dao := service.NewDao()
	defer dao.Close()

	// Register Handler
	user.RegisterUserHandler(srv.Server(), handler.NewUserHandler(srv, dao))

	// Register Struct as Subscriber
	micro.RegisterSubscriber(name, srv.Server(), new(subscriber.User))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
