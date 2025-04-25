package airticle

import (
	"time"

	"github.com/samber/lo"

	"github.com/chaihaobo/boice-blog-api/model/dto/tag"
	"github.com/chaihaobo/boice-blog-api/model/entity"
)

type (
	CreateRequest struct {
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Content     string     `json:"content"`
		Tags        []*tag.Tag `json:"tags"`
	}

	EditRequest struct {
		CreateRequest
	}
)

func (r CreateRequest) ToEntity() *entity.Article {
	return &entity.Article{
		BaseEntity: entity.BaseEntity{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Title:       r.Title,
		Description: r.Description,
		Content:     r.Content,
		Tags: lo.Map(r.Tags, func(tag *tag.Tag, _ int) *entity.Tag {
			return &entity.Tag{
				Name: tag.Name,
				BaseEntity: entity.BaseEntity{
					ID:        tag.ID,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			}
		}),
	}
}
