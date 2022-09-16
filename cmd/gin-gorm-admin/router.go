package main

import (
	v1 "github.com/dot123/gin-gorm-admin/api/v1"
	"github.com/dot123/gin-gorm-admin/internal/middleware"
	"github.com/dot123/gin-gorm-admin/internal/middleware/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var _ IRouter = (*Router)(nil)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

type Router struct {
	MyJwt      *jwt.JWT
	UserApi    *v1.UserApi
	MonitorApi *v1.MonitorApi
	FileApi    *v1.FileApi
	MsgApi     *v1.MsgApi
}

func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

func (a *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}

// RegisterAPI register api group router
func (a *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")
	g.Use(middleware.RateLimiterMiddleware())

	v1 := g.Group("/v1")

	authMiddleware := a.MyJwt.GinJWTMiddlewareInit(new(jwt.AllUserAuthorizator))
	adminMiddleware := a.MyJwt.GinJWTMiddlewareInit(new(jwt.AdminAuthorizator))

	app.NoRoute(authMiddleware.MiddlewareFunc(), middleware.NoRouteHandler())

	v1.POST("/login", authMiddleware.LoginHandler)
	v1.GET("/refreshToken", authMiddleware.RefreshHandler)
	v1.POST("/logout", authMiddleware.LogoutHandler)

	a.FileApi.RegisterRoute(v1.Group("/public"))
	a.UserApi.RegisterRoute(v1.Group("/user"), authMiddleware, adminMiddleware)
	a.MsgApi.RegisterRoute(v1.Group("/msg"), authMiddleware, adminMiddleware)
	a.MonitorApi.RegisterRoute(v1.Group("/monitor"), authMiddleware, adminMiddleware)
}
