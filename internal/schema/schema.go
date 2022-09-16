package schema

import "github.com/dot123/gin-gorm-admin/internal/validate"

type Pagination struct {
	PageNum  int `form:"pageNum,default=1"`
	PageSize int `form:"pageSize,default=10" validate:"max=100"`
}

func (m *Pagination) Verify() string {
	messages := map[string]string{
		"PageSize.max": "请求页数最大为100",
	}

	ok, err := validate.VerifyReturnOneError(m, messages)
	if !ok {
		return err
	}

	return ""
}
