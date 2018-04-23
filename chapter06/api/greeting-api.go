package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	hello "../proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
)

type Say struct {
	Client hello.SayClient
}

func (s *Say) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Say.Hello request - Micro Greeter API")
	name, ok := req.Get["name"]
	if ok {
		response, err := s.Client.Hello(ctx, &hello.Request{
			Name: strings.Join(name.Values, " "),
		})
		if err != nil {
			return err
		}
		message, _ := json.Marshal(map[string]string{
			"message": response.Msg,
		})
		rsp.Body = string(message)
	}
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.greeter"),
	)
	service.Init()
	service.Server().Handle(
		service.Server().NewHandler(
			&Say{Client: hello.NewSayClient("go.micro.service.greeter", service.Client())},
		),
	)
	if err := service.Run(); err != nil {
		log.Fatal("error starting micro api : ", err)
		return
	}
}
