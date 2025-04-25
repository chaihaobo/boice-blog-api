package user

import (
	"github.com/gin-gonic/gin"

	"github.com/chaihaobo/boice-blog-api/application"
	"github.com/chaihaobo/boice-blog-api/constant"
	"github.com/chaihaobo/boice-blog-api/model/dto/user"
	"github.com/chaihaobo/boice-blog-api/resource"
)

type (
	Controller interface {
		Login(ctx *gin.Context) (*user.LoginResponse, error)
		VerifyPermission(ctx *gin.Context) (bool, error)
	}

	controller struct {
		app application.Application
		res resource.Resource
	}
)

func (c *controller) VerifyPermission(ctx *gin.Context) (bool, error) {
	password, ok := ctx.Value(constant.ContextKeyPassword).(string)
	return ok && password == c.res.Configuration().Service.Password, nil
}

func NewController(res resource.Resource, app application.Application) Controller {
	return &controller{
		app: app,
		res: res,
	}
}
