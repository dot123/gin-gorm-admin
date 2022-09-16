package util

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/contextx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Model struct {
	ID        uint64 `gorm:"primaryKey;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	trans, ok := contextx.FromTrans(ctx)
	if ok && !contextx.FromNoTrans(ctx) {
		db, ok := trans.(*gorm.DB)
		if ok {
			if contextx.FromTransLock(ctx) {
				db = db.Clauses(clause.Locking{Strength: "UPDATE"})
			}
			return db
		}
	}

	return defDB
}

func GetDBWithModel(ctx context.Context, defDB *gorm.DB, m interface{}) *gorm.DB {
	return GetDB(ctx, defDB).Model(m)
}

type TransFunc func(context.Context) error

func ExecTrans(ctx context.Context, db *gorm.DB, fn TransFunc) error {
	transModel := &Trans{DB: db}
	return transModel.Exec(ctx, fn)
}

func ExecTransWithLock(ctx context.Context, db *gorm.DB, fn TransFunc) error {
	if !contextx.FromTransLock(ctx) {
		ctx = contextx.NewTransLock(ctx)
	}
	return ExecTrans(ctx, db, fn)
}

// GetPages 分页返回数据
func GetPages(db *gorm.DB, out interface{}, pageNum, pageSize int) (int64, error) {
	var count int64

	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	} else if count == 0 {
		return count, nil
	}

	return count, db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(out).Error
}
