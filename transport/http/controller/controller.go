package controller

import (
	"github.com/chaihaobo/boice-blog-api/application"
	"github.com/chaihaobo/boice-blog-api/resource"
	"github.com/chaihaobo/boice-blog-api/transport/http/controller/article"
	"github.com/chaihaobo/boice-blog-api/transport/http/controller/health"
	"github.com/chaihaobo/boice-blog-api/transport/http/controller/tag"
	"github.com/chaihaobo/boice-blog-api/transport/http/controller/user"
)

type (
	Controller interface {
		Health() health.Controller
		User() user.Controller
		Article() article.Controller
		Tag() tag.Controller
	}

	controllers struct {
		healthController  health.Controller
		userController    user.Controller
		articleController article.Controller
		tagController     tag.Controller
	}
)

func (c *controllers) Tag() tag.Controller {
	return c.tagController
}

func (c *controllers) User() user.Controller {
	return c.userController
}

func (c *controllers) Health() health.Controller {
	return c.healthController
}

func (c *controllers) Article() article.Controller {
	return c.articleController
}
func New(res resource.Resource, app application.Application) Controller {
	return &controllers{
		healthController:  health.NewController(res, app),
		userController:    user.NewController(res, app),
		articleController: article.NewController(res, app),
		tagController:     tag.NewController(res, app),
	}
}
