package role

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var RoleSet = wire.NewSet(wire.Struct(new(RoleRepo), "*"))

type RoleRepo struct {
	DB *gorm.DB
}

func (a *RoleRepo) FindAllByUsername(ctx context.Context, userName string) (*[]*Role, error) {
	roles := new([]*Role)
	err := GetRoleDB(ctx, a.DB).Where("`user_name` = ?", userName).Find(roles).Error
	return roles, err
}

func (a *RoleRepo) GetPage(ctx context.Context, userID uint64) (*[]*Role, error) {
	roles := new([]*Role)
	err := GetRoleDB(ctx, a.DB).Where("`user_id` = ?", userID).Find(roles).Error
	return roles, err
}

func (a *RoleRepo) Create(ctx context.Context, role *Role) error {
	err := GetRoleDB(ctx, a.DB).Create(role).Error
	return err
}

func (a *RoleRepo) Get(ctx context.Context, userID uint64) (*Role, error) {
	role := new(Role)
	err := GetRoleDB(ctx, a.DB).Where("`user_id` = ?", userID).Take(role).Error
	return role, err
}

func (a *RoleRepo) Delete(ctx context.Context, userID uint64) error {
	err := GetRoleDB(ctx, a.DB).Where("`user_id` = ?", userID).Delete(new(Role)).Error
	return err
}

func (a *RoleRepo) Update(ctx context.Context, role *Role) error {
	err := GetRoleDB(ctx, a.DB).Where("`id` = ?", role.ID).Updates(role).Error
	return err
}
