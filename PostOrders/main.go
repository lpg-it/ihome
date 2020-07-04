package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"ihome/PostOrders/handler"
	postorders "ihome/PostOrders/proto/postorders"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.PostOrders"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	postorders.RegisterPostOrdersHandler(service.Server(), new(handler.Server))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.PostOrders", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.PostOrders", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
