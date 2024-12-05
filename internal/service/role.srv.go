package service

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/errors"
	"github.com/dot123/gin-gorm-admin/internal/models/role"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/dot123/gin-gorm-admin/pkg/utils"
	"github.com/google/wire"
)

var RoleSet = wire.NewSet(wire.Struct(new(RoleSrv), "*"))

type RoleSrv struct {
	RoleRepo *role.RoleRepo
}

func (s *RoleSrv) GetUserRoles(ctx context.Context, userName string) (*[]*schema.Role, error) {
	roles, err := s.RoleRepo.FindAllByUsername(ctx, userName)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, errors.NewDefaultResponse("获取用户身份信息失败")
	}

	result := make([]*schema.Role, 0)
	utils.Copy(&result, roles)

	return &result, nil
}
