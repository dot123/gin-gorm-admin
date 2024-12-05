package jwt

import (
	"context"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dot123/gin-gorm-admin/internal/config"
	"github.com/dot123/gin-gorm-admin/internal/contextx"
	"github.com/dot123/gin-gorm-admin/internal/models/user"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/internal/service"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/dot123/gin-gorm-admin/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"log"
	"time"
)

func wrapUserAuthContext(c *gin.Context, userID uint64, userName string) {
	ctx := contextx.NewUserID(c.Request.Context(), userID)
	ctx = contextx.NewUserName(ctx, userName)
	ctx = logger.NewUserIDContext(ctx, userID)
	ctx = logger.NewUserNameContext(ctx, userName)
	c.Request = c.Request.WithContext(ctx)
}

//### 如果是使用Go Module,gin-jwt模块应使用v2
//下载安装，开启Go Module "go env -w GO111MODULE=on",然后执行"go get github.com/appleboy/gin-jwt/v2"
//导入应写成 import "github.com/appleboy/gin-jwt/v2"
//### 如果不是使用Go Module
//下载安装gin-jwt，"go get github.com/appleboy/gin-jwt"
//导入import "github.com/appleboy/gin-jwt"

var JWTSet = wire.NewSet(wire.Struct(new(JWT), "*"))

type JWT struct {
	UserSrv *service.UserSrv
	RoleSrv *service.RoleSrv
}

// GinJWTMiddlewareInit 初始化中间件
func (j *JWT) GinJWTMiddlewareInit(jwtAuthorizator IAuthorizator) (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       config.C.JWTAuth.Realm,
		Key:         []byte(config.C.JWTAuth.Key),
		Timeout:     config.C.JWTAuth.Expired * time.Second,
		MaxRefresh:  config.C.JWTAuth.Expired * time.Second,
		IdentityKey: config.C.JWTAuth.IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*schema.UserRole); ok {
				roles, _ := j.RoleSrv.GetUserRoles(context.TODO(), v.UserName)

				v.UserRoles = *roles

				jsonRole, _ := json.Marshal(v.UserRoles)
				//maps the claims in the JWT
				return jwt.MapClaims{
					"userName":  v.UserName,
					"userID":    v.UserID,
					"userRoles": utils.B2S(jsonRole),
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			roles := jwt.ExtractClaims(c)
			jsonRole := roles["userRoles"].(string)
			userRoles := make([]*schema.Role, 0)
			json.Unmarshal(utils.S2B(jsonRole), &userRoles)

			userName := roles["userName"].(string)
			userID := uint64(roles["userID"].(float64))

			wrapUserAuthContext(c, userID, userName)

			return &schema.UserRole{
				UserName:  userName,
				UserID:    userID,
				UserRoles: userRoles,
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			ctx := c.Request.Context()
			//handles the login logic. On success LoginResponse is called, on failure Unauthorized is called
			loginVars := new(user.User)
			if err := c.ShouldBind(loginVars); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVars.Username
			password := loginVars.Password

			userID, err := j.UserSrv.CheckUser(ctx, username, password)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return &schema.UserRole{
				UserName: username,
				UserID:   userID,
			}, nil
		},
		//receives identity and handles authorization logic
		Authorizator: jwtAuthorizator.HandleAuthorizator,
		//handles unauthorized logic
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code": code,
				"msg":  message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return
}
