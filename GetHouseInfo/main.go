package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"ihome/GetHouseInfo/handler"
	gethouseinfo "ihome/GetHouseInfo/proto/gethouseinfo"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.GetHouseInfo"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	gethouseinfo.RegisterGetHouseInfoHandler(service.Server(), new(handler.Server))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.GetHouseInfo", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.GetHouseInfo", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
