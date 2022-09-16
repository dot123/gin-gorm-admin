package file

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var FileSet = wire.NewSet(wire.Struct(new(FileRepo), "*"))

type FileRepo struct {
	DB *gorm.DB
}

func (m *FileRepo) Upload(ctx context.Context, file *File) error {
	err := GetFileDB(ctx, m.DB).Create(file).Error
	return err
}

func (m *FileRepo) FindFile(ctx context.Context, id uint64) (*File, error) {
	file := new(File)
	err := GetFileDB(ctx, m.DB).Where("`id` = ?", id).Take(file).Error
	return file, err
}

func (m *FileRepo) DeleteFile(ctx context.Context, id uint64) error {
	err := GetFileDB(ctx, m.DB).Where("`id` = ?", id).Unscoped().Delete(new(File)).Error
	return err
}
