package service

import (
	"context"
	"fmt"
	"github.com/dot123/gin-gorm-admin/internal/config"
	"github.com/dot123/gin-gorm-admin/internal/errors"
	"github.com/dot123/gin-gorm-admin/internal/models/msg"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/dot123/gin-gorm-admin/pkg/redisHelper"
	"github.com/go-redis/redis"
	"github.com/jinzhu/copier"
	"golang.org/x/sync/singleflight"
	"time"
)

const (
	NoticesKey = "notices"
)

func NewMsgSrv(msgRepo *msg.MsgRepo) *MsgSrv {
	rc := config.C.Redis
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": rc.Addr,
		},
		Password: rc.Password,
		DB:       0,
	})
	return &MsgSrv{MsgRepo: msgRepo, Ring: ring, g: singleflight.Group{}}
}

type MsgSrv struct {
	MsgRepo *msg.MsgRepo
	g       singleflight.Group
	Ring    *redis.Ring
}

// GetPage 获取公告信息
func (s *MsgSrv) GetPage(ctx context.Context, params *schema.NoticeGetPageReq) (*schema.NoticeGetPageResp, error) {
	data := new(schema.NoticeGetPageResp)

	// 尝试从缓存中取
	key := fmt.Sprintf("%s:%d-%d", NoticesKey, params.PageNum, params.PageSize)

	// 防止缓存击穿
	val, err, _ := s.g.Do(key, func() (interface{}, error) {
		// 尝试从缓存中取
		err := redisHelper.Get(s.Ring, key, data)
		if err != nil {
			if redis.Nil != err {
				return nil, err
			}
		} else {
			// 从缓存中取到了
			return data, nil
		}

		// 从数据库中取
		result, total, err := s.MsgRepo.GetPage(ctx, params.PageNum, params.PageSize)
		if err != nil {
			return nil, err
		}

		copier.Copy(&data.List, result)
		data.Total = total

		// 再放入缓存
		if err = redisHelper.Set(s.Ring, key, data, 60*time.Second); err != nil {
			return nil, err
		}
		return data, nil
	})

	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, errors.NewDefaultResponse("获取公告信息失败")
	}

	return val.(*schema.NoticeGetPageResp), nil
}

// Create 新建公告
func (s *MsgSrv) Create(ctx context.Context, params *schema.NoticeCreateReq) error {
	model := new(msg.Notice)
	copier.Copy(model, params)

	err := s.MsgRepo.Create(ctx, model)
	redisHelper.LikeDeletes(s.Ring, NoticesKey)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("新建公告失败")
	}
	return nil
}

// Update 更新公告
func (s *MsgSrv) Update(ctx context.Context, params *schema.NoticeUpdateReq) error {
	model, err := s.MsgRepo.Get(ctx, params.ID)
	if err != nil {
		return err
	}

	copier.Copy(model, params)

	err = s.MsgRepo.Update(ctx, model)
	redisHelper.LikeDeletes(s.Ring, NoticesKey)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("更新公告失败")
	}
	return nil
}

// Delete 删除公告
func (s *MsgSrv) Delete(ctx context.Context, id uint64) error {
	err := s.MsgRepo.Delete(ctx, id)
	redisHelper.LikeDeletes(s.Ring, NoticesKey)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("删除公告失败")
	}
	return nil
}
