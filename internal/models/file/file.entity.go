package file

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/dot123/gin-gorm-admin/pkg/types"
	"gorm.io/gorm"
)

type File struct {
	ID        uint64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt types.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt types.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	Name      string     `gorm:"column:name"`
	Url       string     `gorm:"column:url"`
	Tag       string     `gorm:"column:tag"`
	Key       string     `gorm:"column:key"`
}

func GetFileDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(File))
}
