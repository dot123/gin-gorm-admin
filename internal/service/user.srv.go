package service

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/errors"
	"github.com/dot123/gin-gorm-admin/internal/models/role"
	"github.com/dot123/gin-gorm-admin/internal/models/user"
	"github.com/dot123/gin-gorm-admin/internal/models/util"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/google/wire"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var UserSet = wire.NewSet(wire.Struct(new(UserSrv), "*"))

type UserSrv struct {
	TransRepo *util.Trans
	UserRepo  *user.UserRepo
	RoleRepo  *role.RoleRepo
}

func (s *UserSrv) CheckUser(ctx context.Context, username, password string) (uint64, error) {
	user, err := s.UserRepo.Get(ctx, username)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return 0, errors.NewDefaultResponse("获取用户角色失败")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		logger.WithContext(ctx).Errorf("用户密码不正确,参数:%s", password)
		return 0, errors.NewDefaultResponse("用户密码不正确")
	}

	return user.ID, nil
}

func (s *UserSrv) GetUserAvatar(ctx context.Context, username string) (string, error) {
	user, err := s.UserRepo.Get(ctx, username)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return "", errors.NewDefaultResponse("获取用户头像失败")
	}
	return user.Avatar, nil
}

func (s *UserSrv) GetRoles(ctx context.Context, username string) (*[]string, error) {
	user, err := s.UserRepo.Get(ctx, username)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, errors.NewDefaultResponse("获取用户角色失败")
	}

	roles, err := s.RoleRepo.GetPage(ctx, user.ID)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, errors.NewDefaultResponse("获取用户角色失败")
	}

	arrRole := new([]string)
	for _, role := range *roles {
		*arrRole = append(*arrRole, role.Value)
	}

	return arrRole, nil
}

func (s *UserSrv) GetPage(ctx context.Context, req *schema.UserGetPageReq) (*schema.UserGetPageReqResult, error) {
	list, total, err := s.UserRepo.GetPage(ctx, req.PageNum, req.PageSize, req.Name)
	if err != nil {
		return nil, errors.NewDefaultResponse("获取用户信息失败")
	}

	result := new(schema.UserGetPageReqResult)
	for _, u := range *list {
		user := new(schema.User)
		user.ID = u.ID
		user.Username = u.Username
		user.Password = u.Password
		user.Avatar = u.Avatar
		user.UserType = schema.GetUserType(u.UserType)
		user.State = schema.GetStatus(u.State)
		user.CreatedAt = u.CreatedAt
		user.UpdatedAt = u.UpdatedAt
		result.List = append(result.List, user)
	}
	result.Total = total
	return result, nil
}

func (s *UserSrv) Create(ctx context.Context, req *schema.UserCreateReq, createdBy string) error {
	user := new(user.User)
	user.Username = req.Username
	user.UserType = req.UserType
	user.Avatar = req.Avatar

	user.CreatedBy = createdBy
	user.State = 1
	if user.Avatar == "" {
		user.Avatar = "https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG"
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.WithContext(ctx).Errorf("密码加密失败:%s", err.Error())
		return err
	}
	user.Password = string(hash)

	// 此处不能使用事务同时创建用户和角色，因为Role表中需要UserId，而UserId需要插入用户数据后才生成，所以不能用事务，否则会报错
	// 用业务逻辑实现事务效果
	if err := s.UserRepo.Create(ctx, user); err != nil {
		return errors.NewDefaultResponse("新建用户失败")
	}
	// 当成功插入User数据后，user为指针地址，可以获取到ID的值。省去了查数据库拿ID的值步骤
	role := new(role.Role)
	role.UserID = user.ID
	role.UserName = user.Username
	role.Value = "test"
	if user.UserType == 1 {
		role.Value = "admin"
	}
	if err := s.RoleRepo.Create(ctx, role); err != nil {
		// 插入role失败后，删除新插入的用户信息，达到事务处理效果
		s.UserRepo.Delete(ctx, user.ID)
		return errors.NewDefaultResponse("新建用户失败")
	}
	return nil
}

func (s *UserSrv) ExistUserByName(ctx context.Context, username string) bool {
	_, err := s.UserRepo.Get(ctx, username)
	//记录不存在错误(RecordNotFound)，返回false
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	//其他类型的错误，写下日志，返回false
	if err != nil {
		logger.WithContext(ctx).Errorf("判断用户名是否已存在失败:%s", err.Error())
		return false
	}
	return true
}

func (s *UserSrv) Update(ctx context.Context, req *schema.UserUpdateReq, modifiedBy string) error {
	user, err := s.UserRepo.FindOneById(ctx, req.ID)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("更新用户失败")
	}

	user.Password = req.Password
	user.ModifiedBy = modifiedBy
	user.UserType = req.UserType
	user.Avatar = req.Avatar
	role, err := s.RoleRepo.Get(ctx, user.ID)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("更新用户失败")
	}

	role.Value = "test"
	if user.UserType == 1 {
		role.Value = "admin"
	}

	err = s.TransRepo.Exec(ctx, func(ctx context.Context) error {
		if err := s.RoleRepo.Update(ctx, role); err != nil {
			return err
		}
		if err := s.UserRepo.Update(ctx, user); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("更新用户失败")
	}

	return nil
}

func (s *UserSrv) Delete(ctx context.Context, id uint64) error {
	user, err := s.UserRepo.FindOneById(ctx, id)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("删除用户失败")
	}

	if user.Username == "admin" {
		err = errors.New("删除用户失败:不能删除admin账号")
		logger.WithContext(ctx).Error(err)
		return errors.NewDefaultResponse("删除用户失败:不能删除admin账号")
	}

	err = s.TransRepo.Exec(ctx, func(ctx context.Context) error {
		if err := s.RoleRepo.Delete(ctx, id); err != nil {
			return err
		}
		if err := s.UserRepo.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("删除用户失败")
	}

	return nil
}
