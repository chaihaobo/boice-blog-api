package tag

import (
	"context"

	"github.com/chaihaobo/boice-blog-api/constant"
	"github.com/chaihaobo/boice-blog-api/infrastructure"
	tagdto "github.com/chaihaobo/boice-blog-api/model/dto/tag"
	"github.com/chaihaobo/boice-blog-api/resource"
)

type (
	Service interface {
		ListTags(ctx context.Context) ([]*tagdto.Tag, error)
	}

	service struct {
		res   resource.Resource
		infra infrastructure.Infrastructure
	}
)

func (s service) ListTags(ctx context.Context) ([]*tagdto.Tag, error) {
	tags, err := s.infra.Store().Repository().Tag().ListTags(ctx)
	if err != nil {
		s.res.Logger().Error(ctx, "list tags failed for query db", err)
		return nil, constant.ErrSystemMalfunction
	}
	return tagdto.NewTags(tags), nil
}

func NewService(res resource.Resource, infra infrastructure.Infrastructure) Service {
	return &service{
		res:   res,
		infra: infra,
	}
}
