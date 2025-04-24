package controller

import (
	"github.com/chaihaobo/boice-blog-api/application"
	"github.com/chaihaobo/boice-blog-api/resource"
	"github.com/chaihaobo/boice-blog-api/transport/grpc/controller/hello"
)

type (
	Controller interface {
		Hello() hello.Controller
	}
	controller struct {
		hello hello.Controller
	}
)

func (c controller) Hello() hello.Controller {
	return c.hello
}

func NewController(res resource.Resource, app application.Application) Controller {
	return &controller{
		hello: hello.NewController(res, app),
	}
}
