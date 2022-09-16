package v1

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dot123/gin-gorm-admin/internal/ginx"
	"github.com/dot123/gin-gorm-admin/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var MonitorApiSet = wire.NewSet(wire.Struct(new(MonitorApi), "*"))

type MonitorApi struct {
	MonitorSrv *service.MonitorSrv
}

// RegisterRoute 注册路由
func (a *MonitorApi) RegisterRoute(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware, adminMiddleware *jwt.GinJWTMiddleware) {
	r.Use(authMiddleware.MiddlewareFunc())
	{
		r.GET("index", a.Index)
	}
}

// @Tags     服务器状态
// @Summary  获取服务器状态
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Success  200 {object} ginx.ResponseData{data=monitor.Server} "成功结果"
// @Failure  500 {object} ginx.ResponseFail{}                    "失败结果"
// @Router   /monitor/index [get]
func (a *MonitorApi) Index(c *gin.Context) {
	ctx := c.Request.Context()
	s, err := a.MonitorSrv.Index(ctx)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, s)
}
