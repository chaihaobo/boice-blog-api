package hello

import (
	"context"

	"github.com/chaihaobo/boice-blog-api/application"
	"github.com/chaihaobo/boice-blog-api/proto"
	"github.com/chaihaobo/boice-blog-api/resource"
)

type (
	Controller interface {
		proto.HelloServiceServer
	}
	controller struct {
		proto.UnimplementedHelloServiceServer
		res resource.Resource
		app application.Application
	}
)

func (c controller) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Reply: "hello: " + request.Name,
	}, nil
}

func NewController(res resource.Resource, app application.Application) Controller {
	return &controller{
		res: res,
		app: app,
	}
}
