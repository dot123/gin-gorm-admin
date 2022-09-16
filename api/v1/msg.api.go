package v1

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dot123/gin-gorm-admin/internal/ginx"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var MsgApiSet = wire.NewSet(wire.Struct(new(MsgApi), "*"))

type MsgApi struct {
	MsgSrv *service.MsgSrv
}

// RegisterRoute 注册路由
func (a *MsgApi) RegisterRoute(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware, adminMiddleware *jwt.GinJWTMiddleware) {
	r.GET("notice", a.GetPage)
	r.Use(adminMiddleware.MiddlewareFunc())
	{
		r.POST("notice", a.Create)
		r.PUT("notice", a.Update)
		r.DELETE("notice/:id", a.Delete)
	}
}

// @Tags     公告管理
// @Summary  获取公告列表
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data query    schema.NoticeGetPageReq                          true "请求参数"
// @Success  200  {object} ginx.ResponseData{data=schema.NoticeGetPageResp} "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}                              "失败结果"
// @Router   /msg/notice [get]
func (a *MsgApi) GetPage(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.NoticeGetPageReq
	if err := ginx.ParseForm(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.MsgSrv.GetPage(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, result)
}

// @Tags     公告管理
// @Summary  新建公告
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.NoticeCreateReq true "请求参数"
// @Success  200  {object} ginx.ResponseData{}    "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}    "失败结果"
// @Router   /msg/notice [post]
func (a *MsgApi) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.NoticeCreateReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	if err := a.MsgSrv.Create(ctx, &req); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     公告管理
// @Summary  更新公告
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.NoticeUpdateReq true "请求参数"
// @Success  200  {object} ginx.ResponseData{}    "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}    "失败结果"
// @Router   /msg/notice [put]
func (a *MsgApi) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.NoticeUpdateReq
	if err := ginx.ParseJSON(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}

	if err := a.MsgSrv.Update(ctx, &req); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}

// @Tags     公告管理
// @Summary  删除公告
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    id  path     uint64              true "公告id"
// @Success  200 {object} ginx.ResponseData{} "成功结果"
// @Failure  500 {object} ginx.ResponseFail{} "失败结果"
// @Router   /msg/notice/{id} [delete]
func (a *MsgApi) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := ginx.ParseParamID(c, "id")
	if err := a.MsgSrv.Delete(ctx, id); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}
