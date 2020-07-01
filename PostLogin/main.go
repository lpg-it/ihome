package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"ihome/PostLogin/handler"

	postlogin "ihome/PostLogin/proto/postlogin"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.PostLogin"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	postlogin.RegisterPostLoginHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.PostLogin", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.PostLogin", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
