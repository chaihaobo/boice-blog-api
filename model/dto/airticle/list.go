package airticle

import (
	"time"

	"github.com/samber/lo"

	"github.com/chaihaobo/boice-blog-api/model/entity"
	"github.com/chaihaobo/boice-blog-api/model/generic"
)

type (
	ListRequest struct {
		generic.Pagination
	}

	Article struct {
		ID          uint64    `json:"id"`
		Title       string    `json:"title"`
		Tags        []string  `json:"tags"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
	}
)

func NewArticles(articles []*entity.Article) []*Article {
	return lo.Map(articles, func(article *entity.Article, _ int) *Article {
		return &Article{
			ID:          article.ID,
			Title:       article.Title,
			Tags:        lo.Map(article.Tags, func(tag *entity.Tag, _ int) string { return tag.Name }),
			Description: article.Description,
			CreatedAt:   article.CreatedAt,
		}
	})
}
