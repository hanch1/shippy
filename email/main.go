package main

import (
	"context"
	"encoding/json"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/broker/nats"
	"github.com/micro/go-plugins/registry/consul"
	"log"
	pb "shippy/user/proto"
)

const topic = "user.created"

var (
	ConsulAddr = "127.0.0.1:8500"
	BrokerAddr = "0.0.0.0:4222"
)

type Subscriber struct{}

func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
	log.Println("Picked up a new message")
	log.Println("Sending email to:", user.Name)
	return nil
}


func main() {

	// 注册服务
	consulReg := consul.NewRegistry(registry.Addrs(ConsulAddr))
	srv := micro.NewService(
		micro.Name("email"),
		micro.Version("latest"),
		micro.Registry(consulReg),
	)
	srv.Init()

	pubSub := nats.NewBroker(broker.Addrs(BrokerAddr))


	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))
	if err := pubSub.Connect(); err != nil {
		log.Fatalf("broker connect error: %v\n", err)
	}

	// 订阅消息
	_, err := pubSub.Subscribe(topic, func(pub broker.Event) error {
		var user *pb.User
		if err := json.Unmarshal(pub.Message().Body, &user); err != nil {
			return err
		}
		log.Printf("[Create User]: %v\n", user)
		go senEmail(user)
		return nil
	})

	if err != nil {
		log.Printf("sub error: %v\n", err)
	}

	if err := srv.Run(); err != nil {
		log.Fatalf("srv run error: %v\n", err)
	}
}

func senEmail(user *pb.User) error {
	log.Println("Picked up a new message")
	log.Println("Sending email to:", user.Name)
	return nil
}