package repository

import (
	"github.com/chaihaobo/boice-blog-api/infrastructure/store/client"
	"github.com/chaihaobo/boice-blog-api/infrastructure/store/repository/article"
	"github.com/chaihaobo/boice-blog-api/infrastructure/store/repository/tag"
	"github.com/chaihaobo/boice-blog-api/infrastructure/store/repository/user"
)

type (
	Repository interface {
		User() user.Repository
		Article() article.Repository
		Tag() tag.Repository
	}
	repository struct {
		userRepository    user.Repository
		articleRepository article.Repository
		tagRepository     tag.Repository
	}
)

func (r *repository) Tag() tag.Repository {
	return r.tagRepository
}

func (r *repository) Article() article.Repository {
	return r.articleRepository
}

func (r *repository) User() user.Repository {
	return r.userRepository
}

func New(client client.Client) Repository {
	return &repository{
		userRepository:    user.NewRepository(client),
		articleRepository: article.NewRepository(client),
		tagRepository:     tag.NewRepository(client),
	}
}
