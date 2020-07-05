package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"ihome/GetUserOrder/handler"
	getuserorder "ihome/GetUserOrder/proto/getuserorder"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.GetUserOrder"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	getuserorder.RegisterGetUserOrderHandler(service.Server(), new(handler.Server))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.GetUserOrder", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.GetUserOrder", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
