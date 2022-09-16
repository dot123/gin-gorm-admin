package ginx

import (
	"github.com/dot123/gin-gorm-admin/internal/errors"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

const (
	prefix     = "gin-gorm-admin"
	ReqBodyKey = prefix + "/req-body"
	ResBodyKey = prefix + "/res-body"
)

// ResponseData 数据返回结构体
type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ResponseFail 返回成功结构体
type ResponseFail struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 验证器
func verify(obj interface{}) error {
	t := reflect.TypeOf(obj)
	m, ok := t.MethodByName("Verify")
	if ok {
		args := []reflect.Value{reflect.ValueOf(obj)}
		resultList := m.Func.Call(args)
		msg := resultList[0].String()
		if msg != "" {
			return errors.New400Response(msg)
		}
	}
	return nil
}

// ResData 数据返回
func ResData(c *gin.Context, data interface{}) {
	resp := ResponseData{
		Code: 200,
		Data: data,
		Msg:  "success",
	}
	ResJSON(c, http.StatusOK, resp)
}

// ResOk 返回操作成功
func ResOk(c *gin.Context) {
	resp := ResponseData{
		Code: 200,
		Msg:  "success",
		Data: "",
	}
	ResJSON(c, http.StatusOK, resp)
}

// ResJSON 返回JSON数据
func ResJSON(c *gin.Context, httpCode int, resp interface{}) {
	c.JSON(httpCode, resp)
	c.Abort()
}

// GetPage 获取每页数量
func GetPage(c *gin.Context) (pageNum, pageSize int) {
	pageNum, _ = strconv.Atoi(c.Query("page"))
	pageSize, _ = strconv.Atoi(c.Query("limit"))
	if pageSize == 0 {
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	return
}

// GetToken Get jwt token from header (Authorization: Bearer xxx)
func GetToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	prefix := "Bearer "
	if token != "" && strings.HasPrefix(token, prefix) {
		token = token[len(prefix):]
	}
	return token
}

// GetBodyData Get body data from context
func GetBodyData(c *gin.Context) []byte {
	if v, ok := c.Get(ReqBodyKey); ok {
		if b, ok := v.([]byte); ok {
			return b
		}
	}
	return nil
}

// ParseParamID Param returns the value of the URL param
func ParseParamID(c *gin.Context, key string) uint64 {
	id, err := strconv.ParseUint(c.Param(key), 10, 64)
	if err != nil {
		return 0
	}
	return id
}

// ParseJSON Parse body json data to struct
func ParseJSON(c *gin.Context, obj interface{}) error {
	c.ShouldBindJSON(obj)
	return verify(obj)
}

// ParseQuery Parse query parameter to struct
func ParseQuery(c *gin.Context, obj interface{}) error {
	c.ShouldBindQuery(obj)
	return verify(obj)
}

// Bind checks the Method and Content-Type to select a binding engine automatically
func Bind(c *gin.Context, obj interface{}) error {
	c.Bind(obj)
	return verify(obj)
}

// ParseForm Parse body form data to struct
func ParseForm(c *gin.Context, obj interface{}) error {
	c.ShouldBindWith(obj, binding.Form)
	return verify(obj)
}

// Response error object and parse error status code
func ResError(c *gin.Context, err error, status ...int) {
	ctx := c.Request.Context()
	var res *errors.ResponseError

	if err != nil {
		if e, ok := err.(*errors.ResponseError); ok {
			res = e
		} else {
			res = errors.UnWrapResponse(errors.ErrInternalServer)
			res.ERR = err
		}
	} else {
		res = errors.UnWrapResponse(errors.ErrInternalServer)
	}

	if len(status) > 0 {
		res.Status = status[0]
	}

	if err := res.ERR; err != nil {
		if res.Message == "" {
			res.Message = err.Error()
		}

		if status := res.Status; status >= 400 && status < 500 {
			logger.WithContext(ctx).Warnf(err.Error())
		} else if status >= 500 {
			logger.WithContext(logger.NewStackContext(ctx, err)).Errorf(err.Error())
		}
	}

	ResJSON(c, res.Status, ResponseFail{Code: res.Code, Msg: res.Message})
}
