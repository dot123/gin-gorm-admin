package user

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var UserSet = wire.NewSet(wire.Struct(new(UserRepo), "*"))

type UserRepo struct {
	DB *gorm.DB
}

func (a *UserRepo) Get(ctx context.Context, username string) (*User, error) {
	user := new(User)
	err := GetUserDB(ctx, a.DB).Where("`username` = ?", username).Take(user).Error
	return user, err
}

func (a *UserRepo) GetPage(ctx context.Context, pageNum int, pageSize int, name string) (*[]*User, int64, error) {
	var total int64
	list := new([]*User)

	db := GetUserDB(ctx, a.DB)
	if name != "" {
		db = db.Where("`username` LIKE ?", "%"+name+"%")
	}

	total, err := util.GetPages(db, list, pageNum, pageSize)
	return list, total, err
}

func (a *UserRepo) Create(ctx context.Context, user *User) error {
	err := GetUserDB(ctx, a.DB).Create(user).Error
	return err
}

func (a *UserRepo) Update(ctx context.Context, user *User) error {
	err := GetUserDB(ctx, a.DB).Where("`id` = ?", user.ID).Updates(user).Error
	return err
}

func (a *UserRepo) Delete(ctx context.Context, id uint64) error {
	err := GetUserDB(ctx, a.DB).Where("`id` = ?", id).Delete(new(User)).Error
	return err
}

func (a *UserRepo) FindOneById(ctx context.Context, id uint64) (*User, error) {
	user := new(User)
	err := GetUserDB(ctx, a.DB).Where("`id` = ?", id).Take(user).Error
	return user, err
}
