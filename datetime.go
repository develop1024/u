package u

import "time"


const (
	Day = time.Hour * 24
	Hour = time.Hour
	Minute = time.Minute
	Second = time.Second
)

// Yesterday 获取昨天日期
func Yesterday() string {
	return TimeFormat(TimeAdd(Day * -1))
}

// Tomorrow 获取明天日期
func Tomorrow() string {
	return TimeFormat(TimeAdd(Day * 1))
}

// TimeFormat time.Time时间格式化
func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// TimeAdd 日期时间加减操作
func TimeAdd(t time.Duration) time.Time {
	return time.Now().Add(t)
}

// Timestamp 获取当前时间时间戳-秒
func Timestamp() int64 {
	return time.Now().Unix()
}

// TimestampMilliSecond 获取当前时间时间戳-毫秒
func TimestampMilliSecond() int64 {
	return time.Now().UnixMilli()
}

// TimestampMicroSecond 获取当前时间时间戳-微秒
func TimestampMicroSecond() int64 {
	return time.Now().UnixMicro()
}

// TimestampNanoSecond 获取当前时间时间戳-纳秒
func TimestampNanoSecond() int64 {
	return time.Now().UnixNano()
}

// DateTime 获取当前日期时间
func DateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Date 获取日期
func Date() string {
	return time.Now().Format("2006-01-02")
}

// Time 获取时间
func Time() string {
	return time.Now().Format("15:04:05")
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
