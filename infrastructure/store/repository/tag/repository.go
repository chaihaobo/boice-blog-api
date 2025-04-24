package tag

import (
	"context"

	"github.com/chaihaobo/boice-blog-api/infrastructure/store/client"
	"github.com/chaihaobo/boice-blog-api/model/entity"
)

type (
	Repository interface {
		ListTags(ctx context.Context) ([]*entity.Tag, error)
	}
	repository struct {
		client client.Client
	}
)

func (r repository) ListTags(ctx context.Context) ([]*entity.Tag, error) {
	var tags = make([]*entity.Tag, 0)
	if err := r.client.DB(ctx).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func NewRepository(client client.Client) Repository {
	return &repository{
		client: client,
	}
}
