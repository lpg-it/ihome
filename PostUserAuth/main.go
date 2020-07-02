package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"ihome/PostUserAuth/handler"
	postuserauth "ihome/PostUserAuth/proto/postuserauth"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.PostUserAuth"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	postuserauth.RegisterPostUserAuthHandler(service.Server(), new(handler.Server))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.PostUserAuth", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.PostUserAuth", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
