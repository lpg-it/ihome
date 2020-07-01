package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"ihome/DeleteSession/handler"
	deletesession "ihome/DeleteSession/proto/deletesession"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.DeleteSession"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	deletesession.RegisterDeleteSessionHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.DeleteSession", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.DeleteSession", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
