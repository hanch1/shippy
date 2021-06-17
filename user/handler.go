package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/broker"
	_ "github.com/micro/go-plugins/broker/nats"
	"golang.org/x/crypto/bcrypt"
	"log"
	"shippy/user/model"
	pb "shippy/user/proto"
)

var (
	topic = "user.created"
)

type handler struct {
	repo         Repository
	tokenService AuthAble
	PubSub       broker.Broker
}

// 实现 user.pb.go 中 UserServiceHandler 定义的方法
// 创建一个新用户时发布一个事件
func (h *handler) Create(context context.Context, req *pb.User, response *pb.Response) error {
	// 对密码进行加密处理
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u := &model.User{
		Id:       req.Id,
		Name:     req.Name,
		Company:  req.Company,
		Email:    req.Email,
		Password: string(password),
	}
	err = h.repo.Create(u)
	if err != nil {
		return err
	}
	response.User = req

	// 发布带有用户所有信息的消息

	if err := h.publishEvent(req); err != nil {
		return err
	}
	return nil
}

func (h *handler) Get(context context.Context, req *pb.User, response *pb.Response) error {
	u, err := h.repo.Get(req.Id)
	if err != nil {
		return err
	}
	log.Printf("[Get]: req id = %s,  resp id = %s\n", req.Id, u.Id)
	response.User = convert(u)
	return nil
}

func (h *handler) GetAll(context context.Context, request *pb.Request, response *pb.Response) error {
	users, err := h.repo.GetAll()
	if err != nil {
		return err
	}
	for _, user := range users {
		response.Users = append(response.Users, convert(user))
	}

	return nil
}

func (h *handler) Auth(context context.Context, req *pb.User, response *pb.Token) error {
	u := convert2(req)
	err := h.repo.GetByEmailAndPassword(u)
	if err != nil {
		log.Printf("[Auth]: find user from repo error : %v\n", err)
		return err
	}
	// 进行密码验证
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return err
	}
	// 生成token
	t, err := h.tokenService.Encode(req)
	if err != nil {
		return err
	}
	response.Token = t
	return nil
}

func (h *handler) ValidateToken(context context.Context, req *pb.Token, resp *pb.Token) error {
	// 验证token是否正确 可以使用redis进行验证
	// 这里假定验证成功
	// TODO
	log.Printf("[ValidateToken]: auth成功, token = %s\n", req.Token)
	resp.Token = req.Token
	return nil
}

// 发送消息通知
func (h *handler) publishEvent(user *pb.User) error {
	body, err := json.Marshal(user)
	if err != nil {
		log.Printf("[publishEvent]: Marshal error")
		return err
	}

	msg := &broker.Message{
		Header: map[string]string{
			"id": user.Id,
		},
		Body: body,
	}
	// 发布 user.created topic 消息
	err = h.PubSub.Publish(topic, msg)
	if err != nil {
		log.Fatalf("[pub] failed: %v\n", err)
	}
	return nil
}

func convert(s *model.User) *pb.User {
	return &pb.User{
		Id:       s.Id,
		Name:     s.Name,
		Company:  s.Company,
		Email:    s.Email,
		Password: s.Password,
	}
}
func convert2(s *pb.User) *model.User {
	return &model.User{
		Id:       s.Id,
		Name:     s.Name,
		Company:  s.Company,
		Email:    s.Email,
		Password: s.Password,
	}
}
