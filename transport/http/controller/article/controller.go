package article

import (
	"github.com/gin-gonic/gin"

	"github.com/chaihaobo/boice-blog-api/application"
	"github.com/chaihaobo/boice-blog-api/constant"
	"github.com/chaihaobo/boice-blog-api/model/dto/airticle"
	"github.com/chaihaobo/boice-blog-api/model/dto/page"
	"github.com/chaihaobo/boice-blog-api/resource"
)

type (
	Controller interface {
		ListArticles(ctx *gin.Context) (*page.Response[*airticle.Article], error)
	}

	controller struct {
		app application.Application
		res resource.Resource
	}
)

func (c controller) ListArticles(ctx *gin.Context) (*page.Response[*airticle.Article], error) {
	var request airticle.ListRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		c.res.Logger().Error(ctx, "list: failed to bind query", err)
		return nil, constant.ErrBadRequest
	}
	return c.app.Article().ListArticles(ctx, &request)
}

func NewController(res resource.Resource, app application.Application) Controller {
	return &controller{
		app: app,
		res: res,
	}
}
