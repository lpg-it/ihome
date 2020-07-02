package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"ihome/PostAvatar/handler"
	postavatar "ihome/PostAvatar/proto/postavatar"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.PostAvatar"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	postavatar.RegisterPostAvatarHandler(service.Server(), new(handler.Server))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.PostAvatar", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.PostAvatar", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
