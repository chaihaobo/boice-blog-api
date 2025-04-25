package airticle

import (
	"time"

	"github.com/samber/lo"

	"github.com/chaihaobo/boice-blog-api/model/dto/tag"
	"github.com/chaihaobo/boice-blog-api/model/entity"
	"github.com/chaihaobo/boice-blog-api/model/generic"
)

type (
	ListRequest struct {
		generic.Pagination
	}

	Article struct {
		ID          uint64     `json:"id"`
		Title       string     `json:"title"`
		Tags        []string   `json:"tags"`
		RawTags     []*tag.Tag `json:"raw_tags"`
		Description string     `json:"description"`
		Content     string     `json:"content"`
		CreatedAt   time.Time  `json:"created_at"`
	}
)

func NewArticles(articles []*entity.Article) []*Article {
	return lo.Map(articles, func(article *entity.Article, _ int) *Article {
		return NewArticle(article, false)
	})
}

func NewArticle(article *entity.Article, withContent bool) *Article {
	return &Article{
		ID:          article.ID,
		Title:       article.Title,
		Tags:        lo.Map(article.Tags, func(tag *entity.Tag, _ int) string { return tag.Name }),
		Description: article.Description,
		Content:     lo.If(withContent, article.Content).Else(""),
		CreatedAt:   article.CreatedAt,
		RawTags: lo.Map(article.Tags, func(t *entity.Tag, _ int) *tag.Tag {
			return &tag.Tag{
				ID:   t.ID,
				Name: t.Name,
			}
		}),
	}
}
