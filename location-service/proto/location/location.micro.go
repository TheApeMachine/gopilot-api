// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/location/location.proto

/*
Package go_micro_srv_location is a generated protocol buffer package.

It is generated from these files:
	proto/location/location.proto

It has these top-level messages:
	Location
	Specification
	Response
*/
package go_micro_srv_location

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for LocationService service

type LocationService interface {
	FindAll(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
}

type locationService struct {
	c    client.Client
	name string
}

func NewLocationService(name string, c client.Client) LocationService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.location"
	}
	return &locationService{
		c:    c,
		name: name,
	}
}

func (c *locationService) FindAll(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "LocationService.FindAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LocationService service

type LocationServiceHandler interface {
	FindAll(context.Context, *Specification, *Response) error
}

func RegisterLocationServiceHandler(s server.Server, hdlr LocationServiceHandler, opts ...server.HandlerOption) error {
	type locationService interface {
		FindAll(ctx context.Context, in *Specification, out *Response) error
	}
	type LocationService struct {
		locationService
	}
	h := &locationServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&LocationService{h}, opts...))
}

type locationServiceHandler struct {
	LocationServiceHandler
}

func (h *locationServiceHandler) FindAll(ctx context.Context, in *Specification, out *Response) error {
	return h.LocationServiceHandler.FindAll(ctx, in, out)
}
