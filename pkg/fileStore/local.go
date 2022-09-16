package fileStore

import (
	"errors"
	"github.com/dot123/gin-gorm-admin/pkg/hash"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

func InitLocal(path string) *Local {
	return &Local{Path: path}
}

type Local struct {
	Path string
}

// UploadFile 上传文件
func (a *Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = hash.MD5([]byte(name))

	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(a.Path, os.ModePerm)
	if mkdirErr != nil {
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := a.Path + "/" + filename

	f, openErr := file.Open() // 读取文件
	if openErr != nil {
		return "", "", errors.New("function file.Open() Filed, err:" + openErr.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return p, filename, nil
}

// DeleteFile 删除文件
func (a *Local) DeleteFile(key string) error {
	p := a.Path + "/" + key
	if err := os.Remove(p); err != nil {
		return errors.New("本地文件删除失败, err:" + err.Error())
	}
	return nil
}
