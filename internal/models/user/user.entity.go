package user

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/dot123/gin-gorm-admin/pkg/types"
	"gorm.io/gorm"
)

type User struct {
	ID         uint64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt  types.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt  types.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	Username   string     `gorm:"column:username"`
	Password   string     `gorm:"column:password"`
	Avatar     string     `gorm:"column:avatar;default:https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG"`
	UserType   int        `gorm:"column:user_type;default:0;NOT NULL"`
	State      int        `gorm:"column:state;default:1;NOT NULL"`
	CreatedBy  string     `gorm:"column:created_by"`
	ModifiedBy string     `gorm:"column:modified_by"`
}

func GetUserDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(User))
}
