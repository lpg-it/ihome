package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"ihome/PutComment/handler"
	putcomment "ihome/PutComment/proto/putcomment"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.PutComment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	putcomment.RegisterPutCommentHandler(service.Server(), new(handler.Server))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.PutComment", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.PutComment", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
