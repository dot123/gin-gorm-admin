package validate

// https://github.com/go-playground/validator/blob/master/README.md#comparisons

import (
	"github.com/dot123/gin-gorm-admin/pkg/types"
	"github.com/dot123/gin-gorm-admin/pkg/validate"
)

// WithValidations 添加验证器
func WithValidations(v validate.Validation) {
	validate.WithValidations(v)
}

// Validate 返回验证器验证结果错误消息 和 bool (是否验证成功)
func Validate(s any, message map[string]string) (bool, map[string]string) {
	return validate.CustomValidator.Verify(s, message)
}

// VerifyReturnOneError 返回验证器验证结果错误消息 和 bool (是否验证成功)
func VerifyReturnOneError(s any, message map[string]string) (bool, string) {
	return validate.CustomValidator.VerifyReturnOneError(s, message)
}

// ValidateMap map 验证器
func ValidateMap(data map[string]any, rules map[string]any, message map[string]string) (bool, map[string]string) {
	return validate.CustomValidator.ValidateMap(data, rules, message)
}

// ValidateMapReturnOneError map 验证器
func ValidateMapReturnOneError(data map[string]any, rules map[string]any, message map[string]string) (bool, string) {
	return validate.CustomValidator.ValidateMapReturnOneError(data, rules, message)
}

// Var 验证器
func Var(data string, rule string) (bool, error) {
	return validate.CustomValidator.Var(data, rule)
}

// ValidateTime 验证时间
func ValidateTime(t types.Time) (bool, string) {
	if t.String() == "0001-01-01 00:00:00" {
		return false, "时间格式错误"
	}
	return true, ""
}
