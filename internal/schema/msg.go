package schema

import (
	"github.com/dot123/gin-gorm-admin/internal/validate"
	"github.com/dot123/gin-gorm-admin/pkg/types"
)

type Notice struct {
	ID        uint64     `json:"id"`
	CreatedAt types.Time `json:"createdAt"`
	StartTime types.Time `json:"startTime"`
	EndTime   types.Time `json:"endTime"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Operator  string     `json:"operator"`
}

type NoticeGetPageReq struct {
	Pagination
}

type NoticeGetPageResp struct {
	List  []*Notice `json:"list"`
	Total int64     `json:"total"`
}

type NoticeCreateReq struct {
	StartTime types.Time `json:"startTime" validate:"required"`
	EndTime   types.Time `json:"endTime" validate:"required"`
	Title     string     `json:"title" validate:"required,min=2,max=150"`
	Content   string     `json:"content" validate:"required,min=2"`
	Operator  string     `json:"operator" validate:"required,min=2,max=150"`
}

func (m *NoticeCreateReq) Verify() string {
	messages := map[string]string{
		"Title.required":    "标题不能为空",
		"Title.min":         "标题最小字符2个",
		"Title.max":         "标题最大字符150个",
		"Content.required":  "内容不能为空",
		"Content.min":       "内容最小字符2个",
		"Operator.required": "操作者不能为空",
		"Operator.min":      "操作者最小字符2个",
		"Operator.max":      "操作者最大字符150个",
	}

	ok, err := validate.ValidateTime(m.StartTime)
	if !ok {
		return err
	}

	ok, err = validate.ValidateTime(m.EndTime)
	if !ok {
		return err
	}

	ok, err = validate.VerifyReturnOneError(m, messages)
	if !ok {
		return err
	}

	return ""
}

type NoticeUpdateReq struct {
	ID        uint64     `json:"id" validate:"required"`
	StartTime types.Time `json:"startTime" validate:"required"`
	EndTime   types.Time `json:"endTime" validate:"required"`
	Title     string     `json:"title" validate:"required,min=2,max=150"`
	Content   string     `json:"content" validate:"required,min=2"`
	Operator  string     `json:"operator" binding:"required,min=2,max=150"`
}

func (m *NoticeUpdateReq) Verify() string {
	messages := map[string]string{
		"ID.required":       "id不能为空",
		"Title.required":    "标题不能为空",
		"Title.min":         "标题最小字符2个",
		"Title.max":         "标题最大字符150个",
		"Content.required":  "内容不能为空",
		"Content.min":       "内容最小字符2个",
		"Operator.required": "操作者不能为空",
		"Operator.min":      "操作者最小字符2个",
		"Operator.max":      "操作者最大字符150个",
	}

	ok, err := validate.ValidateTime(m.StartTime)
	if !ok {
		return err
	}

	ok, err = validate.ValidateTime(m.EndTime)
	if !ok {
		return err
	}

	ok, err = validate.VerifyReturnOneError(m, messages)
	if !ok {
		return err
	}

	return ""
}
