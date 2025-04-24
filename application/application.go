package application

import (
	"github.com/chaihaobo/boice-blog-api/application/article"
	"github.com/chaihaobo/boice-blog-api/application/health"
	"github.com/chaihaobo/boice-blog-api/application/tag"
	"github.com/chaihaobo/boice-blog-api/application/user"
	"github.com/chaihaobo/boice-blog-api/infrastructure"
	"github.com/chaihaobo/boice-blog-api/resource"
)

type (
	Application interface {
		Health() health.Service
		User() user.Service
		Article() article.Service
		Tag() tag.Service
	}

	application struct {
		health  health.Service
		user    user.Service
		article article.Service
		tag     tag.Service
	}
)

func (a *application) Tag() tag.Service {
	return a.tag
}

func (a *application) Article() article.Service {
	return a.article
}

func (a *application) User() user.Service {
	return a.user
}

func (a *application) Health() health.Service {
	return a.health
}

func New(res resource.Resource, infra infrastructure.Infrastructure) Application {
	return &application{
		health:  health.NewService(res, infra),
		user:    user.NewService(res, infra),
		article: article.NewService(res, infra),
		tag:     tag.NewService(res, infra),
	}
}
