package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"ihome/PutOrders/handler"
	putorders "ihome/PutOrders/proto/putorders"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.PutOrders"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	putorders.RegisterPutOrdersHandler(service.Server(), new(handler.Server))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.PutOrders", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.PutOrders", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
