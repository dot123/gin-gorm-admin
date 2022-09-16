package schema

import (
	"github.com/dot123/gin-gorm-admin/internal/validate"
	"github.com/dot123/gin-gorm-admin/pkg/types"
)

var status = map[int]string{
	0: "禁用",
	1: "正常",
}

var userType = map[int]string{
	1: "管理员",
	2: "测试用户",
}

func GetStatus(code int) string {
	if s, ok := status[code]; ok {
		return s
	}
	return status[0]
}

func GetUserType(code int) string {
	if s, ok := userType[code]; ok {
		return s
	}
	return userType[0]
}

type Role struct {
	ID       uint64 `json:"id"`
	UserID   uint64 `json:"user_id"`
	UserName string `json:"user_name"`
	Value    string `json:"value"`
}

type User struct {
	ID        uint64     `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Avatar    string     `json:"avatar"`
	UserType  string     `json:"user_type"`
	State     string     `json:"state"`
	CreatedAt types.Time `json:"createdAt"`
	UpdatedAt types.Time `json:"updatedAt"`
}

type UserInfo struct {
	Roles        []string
	Introduction string
	Avatar       string
	Name         string
}

type UserRole struct {
	UserName  string
	UserID    uint64
	UserRoles []*Role
}

type UserGetPageReq struct {
	Pagination
	Name string `form:"name"`
}

type UserGetPageReqResult struct {
	List  []*User `json:"list"`
	Total int64   `json:"total"`
}

type UserCreateReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6,max=12"`
	UserType int    `json:"user_type" validate:"required"`
	Avatar   string `json:"avatar"`
}

func (m *UserCreateReq) Verify() string {
	messages := map[string]string{
		"Username.required": "用户名不能为空",
		"Password.required": "密码不能为空",
		"Password.min":      "密码最小长度为6个",
		"Password.max":      "密码最大长度为12个",
		"UserType.required": "用户类型不能为空",
	}

	ok, err := validate.VerifyReturnOneError(m, messages)
	if !ok {
		return err
	}

	return ""
}

type UserUpdateReq struct {
	ID       uint64 `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"max=12"`
	UserType int    `json:"user_type" validate:"required"`
	Avatar   string `json:"avatar"`
}

func (m *UserUpdateReq) Verify() string {
	messages := map[string]string{
		"ID.required":       "id不能为空",
		"Username.required": "用户名不能为空",
		"Password.max":      "密码最大长度为12个",
		"UserType.required": "用户类型不能为空",
	}

	ok, err := validate.VerifyReturnOneError(m, messages)
	if !ok {
		return err
	}

	return ""
}
