package main

// 注册并启动服务
import (
	"context"
	"errors"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/registry/consul"
	"log"
	pb "shippy/consignment/proto"
	userPb "shippy/user/proto"
	vesselPb "shippy/vessel/proto"
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

	// 注册服务 go-micro默认是mdns
	consulReg := consul.NewRegistry(registry.Addrs(CONSUL_ADDR))
	server := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("consignment"),
		micro.Version("latest"),
		micro.Registry(consulReg),
		// 用于鉴权
		micro.WrapHandler(AuthWrapper),
	)
	server.Init()

	// 调用vessel微服务
	vClient := vesselPb.NewVesselServiceClient("vessel", server.Client())
	// 注册 consignment 微服务
	pb.RegisterConsignmentServiceHandler(server.Server(), &handler{
		session:      session,
		vesselClient: vClient,
	})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// AuthWrapper 是一个高阶函数，入参是 ”下一步“ 函数，出参是认证函数
// 在返回的函数内部处理完认证逻辑后，再手动调用 fn() 进行下一步处理
// token 是从 consignment-ci 上下文中取出的，再调用 user-service 将其做验证
// 认证通过则 fn() 继续执行，否则报错
func AuthWrapper(fh server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		token := meta["Token"]
		// Auth
		authClient := userPb.NewUserServiceClient("user", client.DefaultClient)
		authResp, err := authClient.ValidateToken(context.Background(), &userPb.Token{
			Token: token,
		})
		log.Println("Auth Resp:", authResp)
		if err != nil {
			return err
		}
		err = fh(ctx, req, rsp)
		return err
	}
}
