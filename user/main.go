package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/broker/nats"
	"github.com/micro/go-plugins/registry/consul"
	_ "github.com/micro/go-plugins/transport/nats"
	"log"
	"os"
	user "shippy/user/proto"
)

var (
	DSN        = "root:root@tcp(127.0.0.1:3306)/shippy?charset=utf8mb4&parseTime=True&loc=Local"
	ConsulAddr = "127.0.0.1:8500"
	BrokerAddr = "0.0.0.0:4222"
)

func main() {
	db, err := ConnectMysql(DSN)
	if err != nil {
		fmt.Printf("mysql connect error: %v\n", err)
		os.Exit(1)
	}
	repo := &UserRepository{db: db}
	t := &TokenService{}
	// 配置NATS
	pubSub := nats.NewBroker(broker.Addrs(BrokerAddr))
	handler := &handler{
		repo:         repo,
		tokenService: t,
		PubSub:       pubSub,
	}

	// 注册服务
	consulReg := consul.NewRegistry(registry.Addrs(ConsulAddr))
	server := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("user"),
		micro.Version("latest"),
		// 服务注册
		micro.Registry(consulReg),
		// 消息中间件
		micro.Broker(pubSub),
	)
	server.Init()

	user.RegisterUserServiceHandler(server.Server(), handler)

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
