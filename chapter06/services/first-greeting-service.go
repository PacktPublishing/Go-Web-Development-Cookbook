package main

import (
	"context"
	"log"
	"time"

	hello "../proto"
	"github.com/micro/go-micro"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Print("Received Say.Hello request - first greeting service")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.service.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()
	hello.RegisterSayHandler(service.Server(), new(Say))
	if err := service.Run(); err != nil {
		log.Fatal("error starting service : ", err)
		return
	}
}
