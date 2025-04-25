package middleware

import (
	"context"
	"net/http"

	"github.com/chaihaobo/gocommon/restkit"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/glob"
	"github.com/samber/lo"

	"github.com/chaihaobo/boice-blog-api/application"
	"github.com/chaihaobo/boice-blog-api/constant"
	"github.com/chaihaobo/boice-blog-api/resource"
)

const (
	headerKeyAuthorization = "Authorization"
)

type (
	httpEndpoint struct {
		Method string
		Path   string
	}
)

func (h *httpEndpoint) Match(request *http.Request) bool {
	return request.Method == h.Method && glob.MustCompile(h.Path).Match(request.URL.Path)
}

var (
	authList = []*httpEndpoint{
		{Method: http.MethodPost, Path: "/articles"},
	}
)

func AuthMiddleware(res resource.Resource, app application.Application) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		token := gctx.GetHeader(headerKeyAuthorization)
		gctx.Request = gctx.Request.WithContext(context.WithValue(gctx.Request.Context(), constant.ContextKeyPassword, token))
		if isWhiteListRequest(gctx.Request) {
			gctx.Next()
			return
		}
		if token == "" {
			restkit.HTTPWriteErr(gctx.Writer, constant.ErrUnauthorized)
			gctx.Abort()
			return
		}

		if res.Configuration().Service.Password != token {
			restkit.HTTPWriteErr(gctx.Writer, constant.ErrUnauthorized)
			gctx.Abort()
		}
		gctx.Next()
	}
}

func isWhiteListRequest(request *http.Request) bool {
	if _, ok := lo.Find(authList, func(endpoint *httpEndpoint) bool {
		return endpoint.Match(request)
	}); ok {
		return false
	}
	return true
}
