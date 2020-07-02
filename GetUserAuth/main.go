package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"ihome/GetUserAuth/handler"
	getuserauth "ihome/GetUserAuth/proto/getuserauth"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.GetUserAuth"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	getuserauth.RegisterGetUserAuthHandler(service.Server(), new(handler.Server))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.GetUserAuth", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.GetUserAuth", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
