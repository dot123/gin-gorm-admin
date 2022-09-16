package v1

import (
	"github.com/dot123/gin-gorm-admin/internal/ginx"
	"github.com/dot123/gin-gorm-admin/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var FileApiSet = wire.NewSet(wire.Struct(new(FileApi), "*"))

type FileApi struct {
	FileSrv *service.FileSrv
}

// RegisterRoute 注册路由
func (a *FileApi) RegisterRoute(r *gin.RouterGroup) {
	r.POST("uploadFile", a.UploadFile)
}

// @Tags    文件管理
// @Summary 上传文件
// @Accept  multipart/form-data
// @Produce application/json
// @Param   file formData file                true "file"
// @Success 200  {object} ginx.ResponseData{} "成功结果"
// @Failure 500  {object} ginx.ResponseFail{} "失败结果"
// @Router  /public/uploadFile [post]
func (a *FileApi) UploadFile(c *gin.Context) {
	ctx := c.Request.Context()

	file, err := c.FormFile("file")
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	url, err := a.FileSrv.UploadFile(ctx, file, "0")
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResData(c, url)
}
