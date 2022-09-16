package v1

import "github.com/google/wire"

var ProviderSet = wire.NewSet(FileApiSet, MonitorApiSet, UserSet, MsgApiSet)
