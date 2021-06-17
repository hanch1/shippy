package main

// 实现微服务的服务端，处理业务逻辑

import (
	"context"
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	pb "shippy/consignment/proto"
	vesselPb "shippy/vessel/proto"
)

// 微服务服务端 struct handler 必须实现 protobuf 中定义的 rpc 方法
// 实现方法的传参 可参考生成的 consignment.pb.go
type handler struct {
	session      *mgo.Session
	vesselClient vesselPb.VesselServiceClient
}

// 从主会话中 Clone() 出新会话处理查询
func (h *handler) GetRepo() Repository {
	// handler.go 的 GetRepo() 中我们使用 Clone() 来创建新的数据库连接
	// 如果每次查询都用主会话，那所有请求都是同一个底层 socket 执行查询，
	// 后边的请求将会阻塞，不能发挥 Go 天生支持并发的优势
	// 为了避免请求的阻塞，mgo 库提供了 Copy() 和 Clone() 函数来创建新会话
	// Clone 出来的新会话重用了主会话的 socket，
	// 避免了创建 socket 在三次握手时间、资源上的开销，尤其适合那些快速写入的请求。
	// 如果进行了复杂查询、大数据量操作时依旧会阻塞 socket 导致后边的请求阻塞。
	// Copy 为会话创建新的 socket，开销大
	return &ConsignmentRepository{h.session.Clone()}
}

func (h *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	defer h.GetRepo().Close()
	fmt.Printf("[test]: %v\n", req)
	// 检查是否有适合的货轮
	vReq := &vesselPb.Specification{
		Capacity:  int32(len(req.Containers)),
		MaxWeight: req.Weight,
	}
	fmt.Printf("[test]: %v\n", vReq)
	// 根据货物寻找可用的货轮
	vResp, err := h.vesselClient.FindAvailable(context.Background(), vReq)
	if err != nil {
		return err
	}
	// 货物被承运
	log.Printf("found vessel: %s\n", vResp.Vessel.Name)
	req.VesselId = vResp.Vessel.Id
	err = h.GetRepo().Create(req)
	if err != nil {
		return err
	}
	resp.Created = true
	resp.Consignment = req
	return nil
}

func (h *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	defer h.GetRepo().Close()
	consignments, err := h.GetRepo().GetAll()
	if err != nil {
		return err
	}
	resp.Consignments = consignments
	return nil
}

//
