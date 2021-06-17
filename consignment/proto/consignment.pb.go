// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consignment.proto

package consignment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 货物
type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{0}
}

func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (m *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(m, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

// 返回消息(托运结果)
type Response struct {
	Created              bool           `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment   `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	Consignments         []*Consignment `protobuf:"bytes,3,rep,name=consignments,proto3" json:"consignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

// 集装箱
type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{2}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// 查看货物信息的请求
// 客户端想要从服务端请求数据，必须有请求格式，哪怕为空
type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{3}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Consignment)(nil), "consignment.Consignment")
	proto.RegisterType((*Response)(nil), "consignment.Response")
	proto.RegisterType((*Container)(nil), "consignment.Container")
	proto.RegisterType((*GetRequest)(nil), "consignment.GetRequest")
}

func init() { proto.RegisterFile("consignment.proto", fileDescriptor_3804bf87090b51a9) }

var fileDescriptor_3804bf87090b51a9 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcf, 0x4f, 0xc2, 0x30,
	0x18, 0x75, 0xfc, 0xde, 0x37, 0x82, 0xe1, 0x4b, 0x84, 0x46, 0x0f, 0x2e, 0x3b, 0x71, 0xc2, 0x04,
	0x13, 0x0f, 0xc6, 0x93, 0x4b, 0x24, 0x5c, 0xeb, 0xcd, 0x8b, 0xc1, 0xf5, 0x0b, 0x36, 0x91, 0x16,
	0xdb, 0x82, 0xff, 0x8d, 0x89, 0x77, 0xff, 0x48, 0x43, 0xc7, 0xa4, 0x68, 0xb8, 0xed, 0xfd, 0xe8,
	0xfa, 0xde, 0xdb, 0xa0, 0x5f, 0x68, 0x65, 0xe5, 0x42, 0x2d, 0x49, 0xb9, 0xf1, 0xca, 0x68, 0xa7,
	0x31, 0x09, 0xa8, 0xec, 0x3b, 0x82, 0x24, 0xdf, 0x63, 0xec, 0x41, 0x4d, 0x0a, 0x16, 0xa5, 0xd1,
	0x28, 0xe6, 0x35, 0x29, 0x30, 0x85, 0x44, 0x90, 0x2d, 0x8c, 0x5c, 0x39, 0xa9, 0x15, 0xab, 0x79,
	0x21, 0xa4, 0x70, 0x00, 0xad, 0x0f, 0x92, 0x8b, 0x57, 0xc7, 0xea, 0x69, 0x34, 0x6a, 0xf2, 0x1d,
	0xc2, 0x1b, 0x80, 0x42, 0x2b, 0x37, 0x97, 0x8a, 0x8c, 0x65, 0x8d, 0xb4, 0x3e, 0x4a, 0x26, 0x83,
	0x71, 0x18, 0x27, 0xaf, 0x64, 0x1e, 0x38, 0xf1, 0x02, 0xe2, 0x0d, 0x59, 0x4b, 0x6f, 0xcf, 0x52,
	0xb0, 0xa6, 0xbf, 0xaf, 0x53, 0x12, 0x33, 0x91, 0x7d, 0x46, 0xd0, 0xe1, 0x64, 0x57, 0x5a, 0x59,
	0x42, 0x06, 0xed, 0xc2, 0xd0, 0xdc, 0x51, 0x19, 0xb8, 0xc3, 0x2b, 0x88, 0xb7, 0x10, 0x96, 0xf4,
	0xa9, 0x93, 0x09, 0xfb, 0x7b, 0x79, 0xf5, 0xcc, 0x43, 0x33, 0xde, 0x41, 0x37, 0x80, 0x96, 0xd5,
	0x7d, 0xf2, 0xe3, 0x87, 0x0f, 0xdc, 0xd9, 0x12, 0xe2, 0xdf, 0x5a, 0xff, 0xc6, 0xbc, 0x84, 0xa4,
	0x58, 0x5b, 0xa7, 0x97, 0x64, 0xb6, 0xe5, 0xca, 0x31, 0xa1, 0xa2, 0x66, 0x62, 0xbb, 0xa5, 0x36,
	0x72, 0x21, 0x95, 0xdf, 0x32, 0xe6, 0x3b, 0x84, 0x43, 0x68, 0xaf, 0x6d, 0x79, 0xa8, 0x51, 0x0a,
	0x5b, 0x38, 0x13, 0x59, 0x17, 0x60, 0x4a, 0x8e, 0xd3, 0xfb, 0x9a, 0xac, 0x9b, 0x7c, 0x45, 0x80,
	0x41, 0xb4, 0x47, 0x32, 0x1b, 0x59, 0x10, 0x3e, 0x40, 0x3f, 0xf7, 0xc3, 0x84, 0x1f, 0xfa, 0x68,
	0xa1, 0xf3, 0xb3, 0x03, 0xa5, 0x5a, 0x3b, 0x3b, 0xc1, 0x1c, 0x4e, 0xa7, 0xe4, 0x02, 0xab, 0xc5,
	0xe1, 0x81, 0x77, 0x1f, 0xe5, 0xe8, 0x4b, 0xee, 0x7b, 0x4f, 0xdd, 0xab, 0x40, 0x7a, 0x69, 0xf9,
	0x9f, 0xf2, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x18, 0x9a, 0xeb, 0xa9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ConsignmentService service

type ConsignmentServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	// 查看托运货物的信息
	GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error)
}

type consignmentServiceClient struct {
	c           client.Client
	serviceName string
}

func NewConsignmentServiceClient(serviceName string, c client.Client) ConsignmentServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "consignment"
	}
	return &consignmentServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *consignmentServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ConsignmentService.CreateConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consignmentServiceClient) GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ConsignmentService.GetConsignments", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConsignmentService service

type ConsignmentServiceHandler interface {
	CreateConsignment(context.Context, *Consignment, *Response) error
	// 查看托运货物的信息
	GetConsignments(context.Context, *GetRequest, *Response) error
}

func RegisterConsignmentServiceHandler(s server.Server, hdlr ConsignmentServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ConsignmentService{hdlr}, opts...))
}

type ConsignmentService struct {
	ConsignmentServiceHandler
}

func (h *ConsignmentService) CreateConsignment(ctx context.Context, in *Consignment, out *Response) error {
	return h.ConsignmentServiceHandler.CreateConsignment(ctx, in, out)
}

func (h *ConsignmentService) GetConsignments(ctx context.Context, in *GetRequest, out *Response) error {
	return h.ConsignmentServiceHandler.GetConsignments(ctx, in, out)
}
