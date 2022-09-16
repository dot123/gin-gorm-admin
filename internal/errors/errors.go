package errors

import (
	"github.com/pkg/errors"
)

// Define alias
var (
	Is           = errors.Is
	New          = errors.New
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	WithStack    = errors.WithStack
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
)

var (
	ErrExistUser       = NewResponse(0, 412, "用户名已存在")
	ErrNotFound        = NewResponse(0, 404, "页面找不到")
	ErrMethodNotAllow  = NewResponse(0, 405, "方法不允许")
	ErrInternalServer  = NewResponse(0, 500, "内部服务器错误")
	ErrTooManyRequests = NewResponse(0, 429, "请求太频繁")
)
