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
		CreateArticle(ctx context.Context, request *airticle.CreateRequest) error
		EditArticle(ctx context.Context, id uint64, request *airticle.EditRequest) error
		GetArticle(ctx context.Context, id uint64) (*airticle.Article, error)
	}

	service struct {
		res   resource.Resource
		infra infrastructure.Infrastructure
	}
)

func (s service) GetArticle(ctx context.Context, id uint64) (*airticle.Article, error) {
	article, err := s.infra.Store().Repository().Article().Get(ctx, id)
	if err != nil {
		s.res.Logger().Error(ctx, "get article failed for query db", err)
		return nil, constant.ErrSystemMalfunction
	}
	return airticle.NewArticle(article, true), nil
}

func (s service) CreateArticle(ctx context.Context, request *airticle.CreateRequest) error {
	ctx, err := s.infra.Store().Client().Begin(ctx)
	if err != nil {
		s.res.Logger().Error(ctx, "create: failed to begin transaction", err)
		return constant.ErrSystemMalfunction
	}
	defer func() {
		_, _ = s.infra.Store().Client().Rollback(ctx)
	}()
	articleEntity := request.ToEntity()
	// 保存文章
	if err := s.infra.Store().Repository().Article().Save(ctx, articleEntity); err != nil {
		s.res.Logger().Error(ctx, "create: failed to create article", err)
		return constant.ErrSystemMalfunction
	}
	if _, err := s.infra.Store().Client().Commit(ctx); err != nil {
		s.res.Logger().Error(ctx, "create: failed to commit transaction", err)
		return constant.ErrSystemMalfunction
	}
	return nil
}

func (s service) EditArticle(ctx context.Context, id uint64, request *airticle.EditRequest) error {
	ctx, err := s.infra.Store().Client().Begin(ctx)
	if err != nil {
		s.res.Logger().Error(ctx, "create: failed to begin transaction", err)
		return constant.ErrSystemMalfunction
	}
	defer func() {
		_, _ = s.infra.Store().Client().Rollback(ctx)
	}()
	articleEntity, err := s.infra.Store().Repository().Article().Get(ctx, id)
	if err != nil {
		s.res.Logger().Error(ctx, "edit: failed to get article", err)
		return constant.ErrSystemMalfunction
	}
	articleEntity.Edit(request.ToEntity())
	if err := s.infra.Store().Repository().Article().Save(ctx, articleEntity); err != nil {
		s.res.Logger().Error(ctx, "edit: failed to edit article", err)
		return constant.ErrSystemMalfunction
	}
	if _, err := s.infra.Store().Client().Commit(ctx); err != nil {
		s.res.Logger().Error(ctx, "create: failed to commit transaction", err)
		return constant.ErrSystemMalfunction
	}
	return nil
}

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
