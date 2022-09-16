package helper

import (
	"time"
)

// Now 当前时间
func Now() *time.Time {
	location, err := time.LoadLocation("Asia/Shanghai") //设置时区
	if err != nil {
		location = time.FixedZone("CST", 8*3600) //替换上海时区
	}
	nowTime := time.Now().In(location)
	return &nowTime
}

// GetBeforeTime 获取n天前的日期
// _day为负则代表取前几天，为正则代表取后几天，0则为今天
func GetBeforeTime(_day int) *time.Time {
	nowTime := Now()

	// 前n天
	beforeTime := nowTime.AddDate(0, 0, _day)

	return &beforeTime
}

// StringToTime 时间字符转为时间
func StringToTime(date string) (*time.Time, error) {
	location, err := time.LoadLocation("Asia/Shanghai") //设置时区
	if err != nil {
		location = time.FixedZone("CST", 8*3600) //替换上海时区
	}
	t, err := time.ParseInLocation("2006-01-02 15:04:05", date, location)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
