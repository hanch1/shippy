package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	pb "shippy/vessel/proto"
)

func main() {
	// 从consul中发现服务
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	service := micro.NewService(micro.Registry(consulReg))
	service.Init()

	// 测试vessel微服务
	// Create
	client := pb.NewVesselServiceClient("vessel", service.Client())
	vessel := &pb.Vessel{
		Id:        "vessel03",
		Capacity:  50,
		MaxWeight: 650000,
		Name:      "vessel03",
		Available: true,
	}

	response, err := client.Create(context.Background(), vessel)
	if err != nil {
		fmt.Printf("Create Vessel Error: %v\n", err)
	} else {
		fmt.Printf("Create Vessel Success, response = %v\n", response)
	}

}
