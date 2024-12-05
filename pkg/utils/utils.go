package utils

import "github.com/jinzhu/copier"

// Copy 复制结构体
func Copy(toValue interface{}, fromValue interface{}) {
	copier.Copy(toValue, fromValue)
}
