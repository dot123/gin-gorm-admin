package service

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/errors"
	"github.com/dot123/gin-gorm-admin/internal/models/file"
	"github.com/dot123/gin-gorm-admin/pkg/fileStore"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/google/wire"
	"mime/multipart"
	"strings"
)

var FileSet = wire.NewSet(wire.Struct(new(FileSrv), "*"))

type FileSrv struct {
	FileRepo *file.FileRepo
	Local    *fileStore.Local
}

// DeleteFile 删除文件记录
func (s *FileSrv) DeleteFile(ctx context.Context, id uint64) error {
	file, err := s.FileRepo.FindFile(ctx, id)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("删除文件失败")
	}

	if err = s.Local.DeleteFile(file.Key); err != nil {
		logger.WithContext(ctx).Errorf("deleteFile error:%s", err.Error())
		return errors.NewDefaultResponse("删除文件失败")
	}

	if err = s.FileRepo.DeleteFile(ctx, file.ID); err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("删除文件失败")
	}

	return nil
}

// UploadFile 上传到本地
func (s *FileSrv) UploadFile(ctx context.Context, header *multipart.FileHeader, noSave string) (string, error) {
	filePath, key, uploadErr := s.Local.UploadFile(header)
	if uploadErr != nil {
		logger.WithContext(ctx).Errorf("upload error:%s", uploadErr.Error())
		return "", errors.NewDefaultResponse("上传文件失败")
	}
	if noSave == "0" {
		list := strings.Split(header.Filename, ".")
		f := file.File{
			Url:  filePath,
			Name: header.Filename,
			Tag:  list[len(list)-1],
			Key:  key,
		}

		if err := s.FileRepo.Upload(ctx, &f); err != nil {
			logger.WithContext(ctx).Errorf("db error:%s", err.Error())
			return "", errors.NewDefaultResponse("上传文件失败")
		}
		return f.Url, nil
	}
	return "", nil
}
