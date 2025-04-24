package transport

import (
	"context"

	"github.com/chaihaobo/boice-blog-api/application"
	"github.com/chaihaobo/boice-blog-api/infrastructure"
	"github.com/chaihaobo/boice-blog-api/resource"
	"github.com/chaihaobo/boice-blog-api/transport/grpc"
	"github.com/chaihaobo/boice-blog-api/transport/http"
)

type (
	Transport interface {
		ServeAll() error
		ShutdownAll() error
		HTTP() http.Transport
		Grpc() grpc.Transport
	}
	transport struct {
		res  resource.Resource
		http http.Transport
		grpc grpc.Transport
	}
)

func (t *transport) Grpc() grpc.Transport {
	return t.grpc
}

func (t *transport) ShutdownAll() error {
	t.grpc.GracefulStop()
	return t.http.Shutdown()
}

func (t *transport) ServeAll() error {

	go func() {
		if err := t.Grpc().Serve(); err != nil {
			t.res.Logger().Error(context.Background(), "failed to serve grpc server", err)
		}
	}()

	return t.HTTP().Serve()
}

func (t *transport) HTTP() http.Transport {
	return t.http
}

func New(res resource.Resource, infra infrastructure.Infrastructure, app application.Application) Transport {
	httpTransport := http.NewTransport(res, infra, app)
	grpcTransport := grpc.NewTransport(res, infra, app)
	return &transport{
		http: httpTransport,
		grpc: grpcTransport,
	}
}
