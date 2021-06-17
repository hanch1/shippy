package main

import (
	"context"
	"errors"
	"gopkg.in/mgo.v2"
	pb "shippy/vessel/proto"
)

type handler struct {
	session *mgo.Session
}

func (h *handler) GetRepo() Repository {
	return &VesselRepository{h.session.Clone()}
}

// 实现微服务的服务端
func (h *handler) Create(ctx context.Context, req *pb.Vessel, resp *pb.Response) error {
	defer h.GetRepo().Close()
	if err := h.GetRepo().Create(req); err != nil {
		return err
	}
	resp.Vessel = req
	resp.Created = true
	return nil
}

func (h *handler) FindAvailable(context context.Context, req *pb.Specification, resp *pb.Response) error {
	defer h.GetRepo().Close()
	vessels := h.GetRepo().GetVessels()
	// 选择一条容量、载重都符合的货轮
	for _, vessel := range vessels {
		if vessel.Capacity >= req.Capacity && vessel.MaxWeight >= req.MaxWeight {
			resp.Vessel = vessel
			return nil
		}
	}
	// 找不到符合要求的货轮
	return errors.New("No Vessel is available")
}
