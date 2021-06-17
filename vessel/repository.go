package main

// 实现数据库的基本 CURD 操作

import (
	"gopkg.in/mgo.v2"
	pb "shippy/vessel/proto"
)

const (
	DB_NAME        = "shippy"
	CON_COLLECTION = "vessels"
)

type Repository interface {
	Create(*pb.Vessel) error
	Close()
	GetVessels() []*pb.Vessel
}

type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) Create(v *pb.Vessel) error {
	return repo.collection().Insert(v)
}

// 关闭连接
func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) GetVessels() []*pb.Vessel {
	var vessels []*pb.Vessel
	repo.collection().Find(nil).All(&vessels)
	return vessels
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(DB_NAME).C(CON_COLLECTION)
}
