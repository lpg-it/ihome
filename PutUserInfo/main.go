package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"ihome/PutUserInfo/handler"
	putuserinfo "ihome/PutUserInfo/proto/putuserinfo"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.PutUserInfo"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	putuserinfo.RegisterPutUserInfoHandler(service.Server(), new(handler.Server))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.PutUserInfo", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.PutUserInfo", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
