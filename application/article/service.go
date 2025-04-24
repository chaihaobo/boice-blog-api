package article

import (
	"context"

	"github.com/chaihaobo/boice-blog-api/constant"
	"github.com/chaihaobo/boice-blog-api/infrastructure"
	"github.com/chaihaobo/boice-blog-api/model/dto/airticle"
	"github.com/chaihaobo/boice-blog-api/model/dto/page"
	"github.com/chaihaobo/boice-blog-api/resource"
)

type (
	Service interface {
		ListArticles(ctx context.Context, request *airticle.ListRequest) (*page.Response[*airticle.Article], error)
	}

	service struct {
		res   resource.Resource
		infra infrastructure.Infrastructure
	}
)

func (s service) ListArticles(ctx context.Context, request *airticle.ListRequest) (*page.Response[*airticle.Article], error) {
	request.SetupDefault()
	articles, total, err := s.infra.Store().Repository().Article().ListArticles(ctx, request.Pagination)
	if err != nil {
		s.res.Logger().Error(ctx, "list: failed to query articles", err)
		return nil, constant.ErrSystemMalfunction
	}
	return page.NewResponse[*airticle.Article](airticle.NewArticles(articles), total), err

}

func NewService(res resource.Resource, infra infrastructure.Infrastructure) Service {
	return &service{
		res:   res,
		infra: infra,
	}
}
