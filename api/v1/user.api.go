package v1

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dot123/gin-gorm-admin/internal/errors"
	"github.com/dot123/gin-gorm-admin/internal/ginx"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/internal/service"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(wire.Struct(new(UserApi), "*"))

type UserApi struct {
	UserSrv *service.UserSrv
}

// RegisterRoute 注册路由
func (a *UserApi) RegisterRoute(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware, adminMiddleware *jwt.GinJWTMiddleware) {
	r.Use(authMiddleware.MiddlewareFunc())
	{
		r.GET("/info", a.GetUserInfo)
	}

	r.Use(adminMiddleware.MiddlewareFunc())
	{
		r.GET("/list", a.GetPage)
		r.POST("", a.Create)
		r.PUT("", a.Update)
		r.DELETE(":id", a.Delete)
	}
}

// @Tags     UserApi
// @Summary  获取用户信息
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Success  200 {object} ginx.ResponseData{data=schema.UserInfo} "成功结果"
// @Failure  500 {object} ginx.ResponseFail{}                     "失败结果"
// @Router   /user/info [get]
func (a *UserApi) GetUserInfo(c *gin.Context) {
	ctx := c.Request.Context()

	claims := jwt.ExtractClaims(c)
	userName := claims["userName"].(string)
	avatar, err := a.UserSrv.GetUserAvatar(ctx, userName)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	arrRole, err := a.UserSrv.GetRoles(ctx, userName)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	data := schema.UserInfo{Roles: *arrRole, Introduction: "", Avatar: avatar, Name: userName}
	ginx.ResData(c, &data)
}

// @Tags     UserApi
// @Summary  获取用户列表
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data query    schema.UserGetPageReq                               true "请求参数"
// @Success  200  {object} ginx.ResponseData{data=schema.UserGetPageReqResult} "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}                                 "失败结果"
// @Router   /user/list [get]
func (a *UserApi) GetPage(c *gin.Context) {
	ctx := c.Request.Context()
	var req schema.UserGetPageReq
	if err := ginx.ParseForm(c, &req); err != nil {
		ginx.ResError(c, err)
		return
	}
	result, err := a.UserSrv.GetPage(ctx, &req)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, result)
}

// @Tags     UserApi
// @Summary  新建用户
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.UserCreateReq true "请求参数"
// @Success  200  {object} ginx.ResponseData{}  "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}  "失败结果"
// @Router   /user [post]
func (a *UserApi) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var params schema.UserCreateReq
	if err := ginx.Bind(c, &params); err != nil {
		logger.WithContext(ctx).Error(err)
		ginx.ResError(c, err)
	} else {
		claims := jwt.ExtractClaims(c)
		createdBy := claims["userName"].(string)

		if !a.UserSrv.ExistUserByName(ctx, params.Username) {
			if err := a.UserSrv.Create(ctx, &params, createdBy); err != nil {
				ginx.ResError(c, err)
			} else {
				ginx.ResOk(c)
			}
		} else {
			ginx.ResError(c, errors.ErrExistUser)
		}
	}
}

// @Tags     UserApi
// @Summary  修改用户
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     schema.UserUpdateReq true "请求参数"
// @Success  200  {object} ginx.ResponseData{}  "成功结果"
// @Failure  500  {object} ginx.ResponseFail{}  "失败结果"
// @Router   /user [put]
func (a *UserApi) Update(c *gin.Context) {
	ctx := c.Request.Context()

	var params schema.UserUpdateReq
	if err := ginx.ParseJSON(c, &params); err != nil {
		logger.WithContext(ctx).Error(err)
		ginx.ResError(c, err)
	} else {
		claims := jwt.ExtractClaims(c)
		modifiedBy := claims["userName"].(string)
		if err := a.UserSrv.Update(ctx, &params, modifiedBy); err != nil {
			ginx.ResError(c, err)
		} else {
			ginx.ResOk(c)
		}
	}
}

// @Tags     UserApi
// @Summary  删除用户
// @Accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    id  path     uint64              true "id"
// @Success  200 {object} ginx.ResponseData{} "成功结果"
// @Failure  500 {object} ginx.ResponseFail{} "失败结果"
// @Router   /user/{id} [delete]
func (a *UserApi) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	id := ginx.ParseParamID(c, "id")
	if err := a.UserSrv.Delete(ctx, id); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOk(c)
}
