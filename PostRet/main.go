package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"ihome/PostRet/handler"

	postret "ihome/PostRet/proto/postret"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.PostRet"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	postret.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.PostRet", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.PostRet", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
