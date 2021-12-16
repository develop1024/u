package u

import "time"

// Timestamp 获取当前时间时间戳
func Timestamp() int64 {
	return time.Now().Unix()
}

// DateTime 获取当前日期时间
func DateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// TimestampToDatetime 时间戳转格式化日期时间
func TimestampToDatetime(unixTime int64) string {
	t := time.Unix(Timestamp(), 0)
	return t.Format("2006-01-02 15:04:05")
}

// DatetimeToTimestamp 日期时间转时间戳
func DatetimeToTimestamp(datetime string) int64 {
	parse, err := time.Parse("2006-01-02 15:04:05", datetime)
	if err != nil {
		return 0
	}
	return parse.Unix()
}