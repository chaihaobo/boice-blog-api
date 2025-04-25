package tag

import (
	"github.com/gin-gonic/gin"

	"github.com/chaihaobo/boice-blog-api/application"
	tagdto "github.com/chaihaobo/boice-blog-api/model/dto/tag"
	"github.com/chaihaobo/boice-blog-api/resource"
)

type (
	Controller interface {
		ListTags(ctx *gin.Context) ([]*tagdto.Tag, error)
	}

	controller struct {
		app application.Application
		res resource.Resource
	}
)

func (c controller) ListTags(ctx *gin.Context) ([]*tagdto.Tag, error) {
	query := ctx.Query("q")
	return c.app.Tag().ListTags(ctx, query)
}

func NewController(res resource.Resource, app application.Application) Controller {
	return &controller{
		app: app,
		res: res,
	}
}
