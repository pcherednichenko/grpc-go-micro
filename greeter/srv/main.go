package main

import (
	"log"
	"time"

	hello "./proto/hello"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"

	"./redis"
	"golang.org/x/net/context"
)

const counter = "counter"

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) (err error) {
	log.Print("Received Say.Hello request")
	client, err := redis.GetConnection()
	if err != nil {
		return
	}
	counter, err := client.Get(counter).Result()
	if err != nil {
		return
	}
	rsp.Msg = "Hello " + req.Name + ". Counter=" + counter
	return
}

// middleware is a handler wrapper - equals middleware
func middleware(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) (err error) {
		log.Printf("[wrapper] server request: %v", req.Method())
		client, err := redis.GetConnection()
		if err != nil {
			return
		}
		client.Incr(counter)
		err = fn(ctx, req, rsp)
		return err
	}
}

func main() {
	service := grpc.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.WrapHandler(middleware),
	)

	client := redis.InitConnection()
	err := client.Set(counter, 0, 0).Err()
	if err != nil {
		panic(err)
	}

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	hello.RegisterSayHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
