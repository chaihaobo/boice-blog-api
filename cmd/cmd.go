package cmd

import (
	"github.com/chaihaobo/boice-blog-api/application"
	"github.com/chaihaobo/boice-blog-api/cmd/cmder"
	"github.com/chaihaobo/boice-blog-api/cmd/core"
	"github.com/chaihaobo/boice-blog-api/infrastructure"
	"github.com/chaihaobo/boice-blog-api/resource"
	"github.com/chaihaobo/boice-blog-api/transport"
)

func Execute() error {
	ctx, err := initialContext()
	if err != nil {
		return err
	}
	return cmder.NewRoot().Command(ctx).Execute()
}

func initialContext() (*core.Context, error) {
	res, err := resource.New("./configuration.yaml")
	if err != nil {
		return nil, err
	}

	infra, err := infrastructure.New(res)
	if err != nil {
		return nil, err
	}
	app := application.New(res, infra)
	tsp := transport.New(res, infra, app)
	ctx := core.NewContext(res, infra, tsp)
	return ctx, nil
}
