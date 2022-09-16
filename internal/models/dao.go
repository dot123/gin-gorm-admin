package models

import (
	"github.com/dot123/gin-gorm-admin/internal/models/file"
	"github.com/dot123/gin-gorm-admin/internal/models/msg"
	"github.com/dot123/gin-gorm-admin/internal/models/role"
	"github.com/dot123/gin-gorm-admin/internal/models/user"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var RepoSet = wire.NewSet(
	util.TransSet,
	role.RoleSet,
	user.UserSet,
	file.FileSet,
	msg.MsgSet,
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		new(role.Role),
		new(user.User),
		new(file.File),
		new(msg.Notice),
	)
	return err
}
