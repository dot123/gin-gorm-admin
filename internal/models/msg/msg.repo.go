package msg

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var MsgSet = wire.NewSet(wire.Struct(new(MsgRepo), "*"))

type MsgRepo struct {
	DB *gorm.DB
}

func (m *MsgRepo) GetPage(ctx context.Context, pageNum, pageSize int) (*[]*Notice, int64, error) {
	list := new([]*Notice)
	total, err := util.GetPages(GetNoticeDB(ctx, m.DB), list, pageNum, pageSize)
	return list, total, err
}

func (m *MsgRepo) Create(ctx context.Context, model *Notice) error {
	err := GetNoticeDB(ctx, m.DB).Create(model).Error
	return err
}

func (m *MsgRepo) Update(ctx context.Context, model *Notice) error {
	err := GetNoticeDB(ctx, m.DB).Where("`id`=?", model.ID).Save(model).Error
	return err
}

func (m *MsgRepo) Delete(ctx context.Context, id uint64) error {
	err := GetNoticeDB(ctx, m.DB).Where("`id`=?", id).Delete(new(Notice)).Error
	return err
}

func (m *MsgRepo) Get(ctx context.Context, id uint64) (*Notice, error) {
	notice := new(Notice)
	err := GetNoticeDB(ctx, m.DB).First(&notice, id).Error
	return notice, err
}
