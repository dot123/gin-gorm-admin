package service

import "github.com/google/wire"

var ProviderSet = wire.NewSet(FileSet, RoleSet, UserSet, MonitorSet, NewMsgSrv)
