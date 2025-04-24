package tag

import (
	"github.com/samber/lo"

	"github.com/chaihaobo/boice-blog-api/model/entity"
)

type (
	Tag struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
)

func NewTags(tags []*entity.Tag) []*Tag {
	return lo.Map(tags, func(tag *entity.Tag, _ int) *Tag {
		return &Tag{
			ID:   tag.ID,
			Name: tag.Name,
		}
	})
}
