package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"log"
	pb "shippy/vessel/proto"
)

const (
	MONGO_HOST  = "127.0.0.1:27017"
	CONSUL_ADDR = "127.0.0.1:8500"
)

func main() {
	// 创建mongo session
	session, err := CreateSession(MONGO_HOST)
	if err != nil {
		log.Fatalf("create mongoDB session error: %v\n", err)
	}
	defer session.Close()

	// consul
	reg := consul.NewRegistry(registry.Addrs(CONSUL_ADDR))
	server := micro.NewService(
		micro.Name("vessel"),
		micro.Version("latest"),
		micro.Registry(reg),
	)
	server.Init()

	// 注册 vessel 微服务
	pb.RegisterVesselServiceHandler(server.Server(), &handler{
		session: session,
	})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
