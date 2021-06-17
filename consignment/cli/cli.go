package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"io/ioutil"
	"log"
	"os"
	pb "shippy/consignment/proto"
)

const (
	CONSIGNMENT_FILE = "consignment\\cli\\consignment.json"
)

func main() {
	// 从consul中发现服务
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	service := micro.NewService(micro.Registry(consulReg))
	service.Init()



	// 测试consignment微服务
	client := pb.NewConsignmentServiceClient("consignment", service.Client())

	// 在命令行中指定新的货物信息 json 文件
	infoFile := CONSIGNMENT_FILE
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}

	token := "$10$5/52uIBC2ML4Cwbx9/E/7.7TODI/f8hR7TqKrQh0pokMxhG7DMPie"
	// 创建带有用户 token 的 context
	// consignment-service 服务端将从中取出 token，解密取出用户身份
	tokenContext := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	// 解析货物信息
	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v", err)
	} else {
		fmt.Printf("货物：%v\n", consignment)
	}

	// 调用 RPC
	// 将货物存储到我们自己的仓库里
	resp, err := client.CreateConsignment(tokenContext, consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v", err)
	}

	// 新货物是否托运成功
	log.Printf("created: %t", resp.Created)


	// 列出目前所有托运的货物
	resp, err = client.GetConsignments(tokenContext, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("failed to list consignments: %v", err)
	}
	for _, c := range resp.Consignments {
		log.Printf("%+v", c)
	}
}

// 读取 consignment.json 中记录的货物信息
func parseFile(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var consignment *pb.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}
	return consignment, nil
}
