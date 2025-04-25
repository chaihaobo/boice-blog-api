package http

import (
	"context"
	"fmt"
	"net/http"

	ginmiddewate "github.com/chaihaobo/gocommon/middleware/http/gin"
	"github.com/chaihaobo/gocommon/restkit"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"go.uber.org/zap"

	"github.com/chaihaobo/boice-blog-api/application"
	"github.com/chaihaobo/boice-blog-api/infrastructure"
	"github.com/chaihaobo/boice-blog-api/infrastructure/discovery"
	"github.com/chaihaobo/boice-blog-api/model/dto/airticle"
	"github.com/chaihaobo/boice-blog-api/model/dto/page"
	tagdto "github.com/chaihaobo/boice-blog-api/model/dto/tag"
	"github.com/chaihaobo/boice-blog-api/model/dto/user"
	"github.com/chaihaobo/boice-blog-api/resource"
	"github.com/chaihaobo/boice-blog-api/transport/http/controller"
	"github.com/chaihaobo/boice-blog-api/transport/http/middleware"
	"github.com/chaihaobo/boice-blog-api/utils"
)

type (
	Transport interface {
		Serve() error
		Shutdown() error
	}

	transport struct {
		resource   resource.Resource
		engine     *gin.Engine
		infra      infrastructure.Infrastructure
		controller controller.Controller
		server     *http.Server
		serviceID  string
	}
)

func (t *transport) Serve() error {
	var (
		port = t.resource.Configuration().Service.HTTPPort
		name = t.resource.Configuration().Service.Name
	)
	ip, err := utils.GetOutboundIP()
	if err != nil {
		return err
	}
	ctx := context.TODO()
	serviceID, err := t.infra.DiscoveryClient().RegisterService(ctx, &discovery.Service{
		Name:            fmt.Sprintf("%s-http", name),
		IP:              ip,
		Port:            port,
		Type:            discovery.ServiceTypeHTTP,
		HealthCheckPath: "/health",
	})
	if err != nil {
		return fmt.Errorf("failed to register service to consul: %w", err)
	}
	t.serviceID = serviceID
	t.resource.Logger().Info(ctx, "http server started",
		zap.String("name", name),
		zap.Int("port", port))
	return t.server.ListenAndServe()
}

func (t *transport) Shutdown() error {
	ctx := context.TODO()
	if err := t.infra.DiscoveryClient().DeregisterService(ctx, t.serviceID); err != nil {
		t.resource.Logger().Error(ctx, "failed to deregister http service from consul", err)
	}
	return t.server.Shutdown(ctx)
}

func (t *transport) applyRoutes() {
	router := t.engine
	healthController := t.controller.Health()
	userController := t.controller.User()
	articleController := t.controller.Article()
	tagController := t.controller.Tag()

	router.GET("/health", restkit.AdaptToGinHandler(restkit.HandlerFunc[any](healthController.Health)))

	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", restkit.AdaptToGinHandler(restkit.HandlerFunc[*user.LoginResponse](userController.Login)))
		userGroup.POST("/verify-permission", restkit.AdaptToGinHandler(restkit.HandlerFunc[bool](userController.VerifyPermission)))

	}

	articleGroup := router.Group("/articles")
	{
		articleGroup.GET("", restkit.AdaptToGinHandler(restkit.HandlerFunc[*page.Response[*airticle.Article]](articleController.ListArticles)))
		articleGroup.POST("", restkit.AdaptToGinHandler(restkit.HandlerFunc[any](articleController.CreateArticle)))
		articleGroup.GET("/:id", restkit.AdaptToGinHandler(restkit.HandlerFunc[*airticle.Article](articleController.GetArticle)))
		articleGroup.PUT("/:id", restkit.AdaptToGinHandler(restkit.HandlerFunc[any](articleController.EditArticle)))
	}

	tagGroup := router.Group("/tags")
	{
		tagGroup.GET("", restkit.AdaptToGinHandler(restkit.HandlerFunc[[]*tagdto.Tag](tagController.ListTags)))
	}

}

func NewTransport(res resource.Resource, infra infrastructure.Infrastructure, app application.Application) Transport {
	svcConfig := res.Configuration().Service
	gin.SetMode(lo.If(svcConfig.Debug, gin.DebugMode).
		Else(gin.ReleaseMode))
	engine := gin.New()
	engine.Use(
		ginmiddewate.TelemetryMiddleware(svcConfig.Name, svcConfig.Env, res.Logger()),
		func(c *gin.Context) {

			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173,https:/www.chaihaobo.tech,https://boice-blog.onrender.com")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			if c.Request.Method != http.MethodOptions {
				c.Next()
				return
			}
			c.AbortWithStatus(200)
		},
		gin.Recovery(),
		middleware.AuthMiddleware(res, app),
	)
	engine.ContextWithFallback = true

	tsp := &transport{
		resource:   res,
		engine:     engine,
		infra:      infra,
		controller: controller.New(res, app),
		server:     &http.Server{Addr: fmt.Sprintf(":%d", svcConfig.HTTPPort), Handler: engine},
	}
	tsp.applyRoutes()
	return tsp
}
