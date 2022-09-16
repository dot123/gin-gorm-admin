package msg

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/dot123/gin-gorm-admin/pkg/types"
	"gorm.io/gorm"
)

type Notice struct {
	ID        uint64     `gorm:"primary_key;AUTO_INCREMENT;NOT NULL;"`
	CreatedAt types.Time `gorm:"column:created_at;type:dateTime;comment:'创建时间';"`
	StartTime types.Time `gorm:"column:start_time;type:dateTime;comment:'开始时间';"`
	EndTime   types.Time `gorm:"column:end_time;type:dateTime;comment:'结束时间';"`
	Title     string     `gorm:"column:title;comment:'标题';"`
	Content   string     `gorm:"column:content;comment:'内容';"`
	Operator  string     `gorm:"column:operator;comment:'操作者';"`
}

func GetNoticeDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Notice))
}
