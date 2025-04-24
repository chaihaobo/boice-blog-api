package core

import (
	"github.com/chaihaobo/boice-blog-api/infrastructure"
	"github.com/chaihaobo/boice-blog-api/resource"
	"github.com/chaihaobo/boice-blog-api/transport"
)

type Context struct {
	Resource       resource.Resource
	Infrastructure infrastructure.Infrastructure
	Transport      transport.Transport
}

func NewContext(res resource.Resource, infra infrastructure.Infrastructure, tsp transport.Transport) *Context {
	return &Context{
		Resource:       res,
		Infrastructure: infra,
		Transport:      tsp,
	}
}
