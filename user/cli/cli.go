package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	user "shippy/user/proto"
)

func main() {
	// 从consul中发现服务
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	service := micro.NewService(micro.Registry(consulReg))
	service.Init()

	// 测试user微服务
	client := user.NewUserServiceClient("user", service.Client())
	ctx := context.Background()

	// Create
	u := &user.User{
		Id:       "06",
		Name:     "hanci4",
		Company:  "xxx",
		Email:    "24ssd0812605@qq.com",
		Password: "123456",
	}
	response, err := client.Create(ctx, u)
	if err != nil {
		fmt.Printf("create user error: %v\n", err)
	} else {
		fmt.Printf("create user success, resp : %v\n", response)
	}

	// Get
	//u = &user.User{
	//	Id: "01",
	//}
	//res, err := client.Get(ctx, u)
	//if err != nil {
	//	fmt.Printf("get user error: %v\n", err)
	//} else {
	//	fmt.Printf("get user success, resp : %v\n", res)
	//}
	//
	//// GetAll
	//req := &user.Request{}
	//resp, err := client.GetAll(ctx, req)
	//if err != nil {
	//	fmt.Printf("get all user error: %v\n", err)
	//} else {
	//	for _, u := range resp.Users {
	//		fmt.Printf("user: %v\n", u)
	//	}
	//}
	//
	//// Auth
	//u1 := &user.User{
	//	Email:    "hcz@qq.com",
	//	Password: "123456",
	//}
	//token, err := client.Auth(ctx, u1)
	//if err != nil {
	//	fmt.Printf("auth error: %v\n", err)
	//} else {
	//	fmt.Printf("auth success, token = %s\n", token.Token)
	//}

}
