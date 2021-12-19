package u

import (
	"github.com/golang-module/carbon/v2"
	"sync"
	"time"
)

type uTime struct {
	time.Time
}

var _uTime *uTime
var _uTimeOnce sync.Once

// Time 获取uTime对象
func Time() *uTime {
	_uTimeOnce.Do(func() {
		_uTime = &uTime{}
	})
	return _uTime
}

// Now 获取当前时间UTime对象
func (receiver *uTime) Now() *uTime {
	return &uTime{time.Now()}
}

// ParseTime 字符串是时间解析为time.Time时间对象
func (receiver *uTime) ParseTime(t string) time.Time {
	parseTime, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		return time.Time{}
	}
	return parseTime
}

// ParseUTime 字符串是时间解析为UTime时间对象
func (receiver *uTime) ParseUTime(t string) *uTime {
	parseTime, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		return &uTime{}
	}
	return &uTime{parseTime}
}

// TimestampToTimeObject 时间戳转time.Time对象
func (receiver *uTime) TimestampToTimeObject(t int64) time.Time {
	return receiver.ParseTime(receiver.TimestampToDatetime(t))
}

// TimestampToUTimeObject 时间戳转UTime对象
func (receiver *uTime) TimestampToUTimeObject(t int64) *uTime {
	return &uTime{receiver.ParseTime(receiver.TimestampToDatetime(t))}
}

// DateString 获取日期字符串
func (receiver *uTime) DateString() string {
	return receiver.Format("2006-01-02")
}

// TimeString 获取时间字符串
func (receiver *uTime) TimeString() string {
	return receiver.Format("15:04:05")
}

// DateTimeString 获取日期时间字符串
func (receiver *uTime) DateTimeString() string {
	return receiver.Format("2006-01-02 15:04:05")
}

// AddYear 增加1年
func (receiver *uTime) AddYear() *uTime {
	return &uTime{carbon.Now().AddYear().Carbon2Time()}
}

// AddYears 增加年数
func (receiver *uTime) AddYears(years int) *uTime {
	return &uTime{carbon.Now().AddYears(years).Carbon2Time()}
}

// AddMonth 增加一月
func (receiver *uTime) AddMonth() *uTime {
	return &uTime{carbon.Now().AddMonth().Carbon2Time()}
}

// AddMonths 增加月数
func (receiver *uTime) AddMonths(months int) *uTime {
	return &uTime{carbon.Now().AddMonths(months).Carbon2Time()}
}

// AddDay 增加1天
func (receiver *uTime) AddDay() *uTime {
	return &uTime{carbon.Now().AddDay().Carbon2Time()}
}

// AddDays 增加天数
func (receiver *uTime) AddDays(days int) *uTime {
	return &uTime{carbon.Now().AddDays(days).Carbon2Time()}
}

// AddHour 增加1小时
func (receiver *uTime) AddHour() *uTime {
	return &uTime{carbon.Now().AddHour().Carbon2Time()}
}

// AddHours 增加小时数
func (receiver *uTime) AddHours(hours int) *uTime {
	return &uTime{carbon.Now().AddHours(hours).Carbon2Time()}
}

// AddMinute 增加1分钟
func (receiver *uTime) AddMinute() *uTime {
	return &uTime{carbon.Now().AddMinute().Carbon2Time()}
}

// AddMinutes 增加分钟数
func (receiver *uTime) AddMinutes(minutes int) *uTime {
	return &uTime{carbon.Now().AddMinutes(minutes).Carbon2Time()}
}

// AddSecond 增加1秒
func (receiver *uTime) AddSecond() *uTime {
	return &uTime{carbon.Now().AddSecond().Carbon2Time()}
}

// AddSeconds 增加秒数
func (receiver *uTime) AddSeconds(seconds int) *uTime {
	return &uTime{carbon.Now().AddSeconds(seconds).Carbon2Time()}
}

// SubYear 减去1年
func (receiver *uTime) SubYear() *uTime {
	return &uTime{carbon.Now().SubYear().Carbon2Time()}
}

// SubYears 减去年数
func (receiver *uTime) SubYears(years int) *uTime {
	return &uTime{carbon.Now().SubYears(years).Carbon2Time()}
}

// SubMonth 减去1月
func (receiver *uTime) SubMonth() *uTime {
	return &uTime{carbon.Now().SubMonth().Carbon2Time()}
}

// SubMonths 减去月数
func (receiver *uTime) SubMonths(months int) *uTime {
	return &uTime{carbon.Now().SubMonths(months).Carbon2Time()}
}

// SubDay 减去1天
func (receiver *uTime) SubDay() *uTime {
	return &uTime{carbon.Now().SubDay().Carbon2Time()}
}

// SubDays 减去天数
func (receiver *uTime) SubDays(days int) *uTime {
	return &uTime{carbon.Now().SubDays(days).Carbon2Time()}
}

// SubHour 减去1小时
func (receiver *uTime) SubHour() *uTime {
	return &uTime{carbon.Now().SubHour().Carbon2Time()}
}

// SubHours 减去小时数
func (receiver *uTime) SubHours(hours int) *uTime {
	return &uTime{carbon.Now().SubHours(hours).Carbon2Time()}
}

// SubMinute 减去1分钟
func (receiver *uTime) SubMinute() *uTime {
	return &uTime{carbon.Now().SubMinute().Carbon2Time()}
}

// SubMinutes 减去分钟数
func (receiver *uTime) SubMinutes(minutes int) *uTime {
	return &uTime{carbon.Now().SubMinutes(minutes).Carbon2Time()}
}

// SubSecond 减去1秒钟
func (receiver *uTime) SubSecond() *uTime {
	return &uTime{carbon.Now().SubSecond().Carbon2Time()}
}

// SubSeconds 减去秒钟数
func (receiver *uTime) SubSeconds(seconds int) *uTime {
	return &uTime{carbon.Now().SubSeconds(seconds).Carbon2Time()}
}

// ToTime UTime对象转time.Time对象
func (receiver *uTime) ToTime() time.Time {
	return receiver.Time
}

// Yesterday 获取昨天日期
func (receiver *uTime) Yesterday() *uTime {
	return receiver.SubDay()
}

// Tomorrow 获取明天日期
func (receiver *uTime) Tomorrow() *uTime {
	return receiver.AddDay()
}

// TimeFormat time.Time时间格式化
func (receiver *uTime) TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// Timestamp 获取当前时间时间戳-秒
func (receiver *uTime) Timestamp() int64 {
	return time.Now().Unix()
}

// TimestampMilliSecond 获取当前时间时间戳-毫秒
func (receiver *uTime) TimestampMilliSecond() int64 {
	return time.Now().UnixMilli()
}

// TimestampMicroSecond 获取当前时间时间戳-微秒
func (receiver *uTime) TimestampMicroSecond() int64 {
	return time.Now().UnixMicro()
}

// TimestampNanoSecond 获取当前时间时间戳-纳秒
func (receiver *uTime) TimestampNanoSecond() int64 {
	return time.Now().UnixNano()
}

// DateTime 获取当前日期时间
func (receiver *uTime) DateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Date 获取日期
func (receiver *uTime) Date() string {
	return time.Now().Format("2006-01-02")
}

// TimestampToDatetime 时间戳转格式化日期时间
func (receiver *uTime) TimestampToDatetime(unixTime int64) string {
	t := time.Unix(unixTime, 0)
	return t.Format("2006-01-02 15:04:05")
}

// DatetimeToTimestamp 日期时间转时间戳
func (receiver *uTime) DatetimeToTimestamp(datetime string) int64 {
	parse, err := time.Parse("2006-01-02 15:04:05", datetime)
	if err != nil {
		return 0
	}
	return parse.Unix()
}
