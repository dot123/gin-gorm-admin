package service

import (
	"context"
	"github.com/dot123/gin-gorm-admin/internal/errors"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/dot123/gin-gorm-admin/pkg/monitor"
	"github.com/google/wire"
)

var MonitorSet = wire.NewSet(wire.Struct(new(MonitorSrv), "*"))

type MonitorSrv struct{}

func (s *MonitorSrv) Index(ctx context.Context) (*monitor.Server, error) {
	var err error
	var m monitor.Server

	m.Os = monitor.GetOSInfo()
	if m.Cpu, err = monitor.GetCpuInfo(); err != nil {
		logger.WithContext(ctx).Errorf("GetCpuInfo error: %s", err.Error())
		return nil, errors.NewDefaultResponse("获取服务器状态失败")
	}

	if m.Rrm, err = monitor.GetMemInfo(); err != nil {
		logger.WithContext(ctx).Errorf("GetMemInfo error: %s", err.Error())
		return nil, errors.NewDefaultResponse("获取服务器状态失败")
	}

	if m.Disk, err = monitor.GetDiskInfo(); err != nil {
		logger.WithContext(ctx).Errorf("GetDiskInfo error: %s", err.Error())
		return nil, errors.NewDefaultResponse("获取服务器状态失败")
	}

	return &m, nil
}
