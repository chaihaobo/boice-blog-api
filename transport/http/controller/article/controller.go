package article

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"

	"github.com/chaihaobo/boice-blog-api/application"
	"github.com/chaihaobo/boice-blog-api/constant"
	"github.com/chaihaobo/boice-blog-api/model/dto/airticle"
	"github.com/chaihaobo/boice-blog-api/model/dto/page"
	"github.com/chaihaobo/boice-blog-api/resource"
)

type (
	Controller interface {
		ListArticles(ctx *gin.Context) (*page.Response[*airticle.Article], error)
		CreateArticle(ctx *gin.Context) (any, error)
		GetArticle(ctx *gin.Context) (*airticle.Article, error)
		EditArticle(ctx *gin.Context) (any, error)
	}

	controller struct {
		app application.Application
		res resource.Resource
	}
)

func (c controller) GetArticle(ctx *gin.Context) (*airticle.Article, error) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		return nil, constant.ErrBadRequest
	}
	return c.app.Article().GetArticle(ctx, uint64(lo.Must(strconv.Atoi(id))))
}

func (c controller) CreateArticle(ctx *gin.Context) (any, error) {
	var request airticle.CreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}
	return nil, c.app.Article().CreateArticle(ctx, &request)

}

func (c controller) EditArticle(ctx *gin.Context) (any, error) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		return nil, constant.ErrBadRequest
	}
	var request airticle.EditRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}
	return nil, c.app.Article().EditArticle(ctx, uint64(lo.Must(strconv.Atoi(id))), &request)

}

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
