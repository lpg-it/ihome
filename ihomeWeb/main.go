package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-log"
	"net/http"

	"github.com/micro/go-web"
	_ "ihome/ihomeWeb/models"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.ihomeWeb"),
		web.Version("latest"),
		web.Address(":8080"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	rou := httprouter.New()
	rou.NotFound = http.FileServer(http.Dir("html"))

	// register html handler
	service.Handle("/", rou)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
