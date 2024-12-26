package utils

import (
	"time"
)

func ToTimeString(t uint) string {
	return time.Unix(int64(t), 0).Format("2006-01-02 15:04:05")
}

// 要自己确保时间格式是正确的
func StringToTimeUnix(timeStr string) int64 {
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, timeStr)
	return t.Unix()
}
