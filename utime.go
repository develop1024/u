package u

import (
	"github.com/golang-module/carbon/v2"
	"time"
)


type UTime struct{
	time.Time
}

// Now 获取当前时间UTime对象
func Now() UTime {
	return UTime{time.Now()}
}

// ParseTime 字符串是时间解析为time.Time时间对象
func ParseTime(t string) time.Time {
	parseTime, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		return time.Time{}
	}
	return parseTime
}

// ParseUTime 字符串是时间解析为UTime时间对象
func ParseUTime(t string) UTime {
	parseTime, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		return UTime{}
	}
	return UTime{parseTime}
}

// TimestampToTimeObject 时间戳转time.Time对象
func TimestampToTimeObject(t int64) time.Time {
	return ParseTime(TimestampToDatetime(t))
}

// TimestampToUTimeObject 时间戳转UTime对象
func TimestampToUTimeObject(t int64) UTime {
	return UTime{ParseTime(TimestampToDatetime(t))}
}

// DateString 获取日期字符串
func (receiver UTime) DateString() string {
	return receiver.Format("2006-01-02")
}

// TimeString 获取时间字符串
func (receiver UTime) TimeString() string {
	return receiver.Format("15:04:05")
}

// DateTimeString 获取日期时间字符串
func (receiver UTime) DateTimeString() string {
	return receiver.Format("2006-01-02 15:04:05")
}

// AddYear 增加1年
func (receiver UTime) AddYear() UTime {
	return UTime{carbon.Now().AddYear().Carbon2Time()}
}

// AddYears 增加年数
func (receiver UTime) AddYears(years int) UTime {
	return UTime{carbon.Now().AddYears(years).Carbon2Time()}
}

// AddMonth 增加一月
func (receiver UTime) AddMonth() UTime {
	return UTime{carbon.Now().AddMonth().Carbon2Time()}
}

// AddMonths 增加月数
func (receiver UTime) AddMonths(months int) UTime {
	return UTime{carbon.Now().AddMonths(months).Carbon2Time()}
}

// AddDay 增加1天
func (receiver UTime) AddDay() UTime {
	return UTime{carbon.Now().AddDay().Carbon2Time()}
}

// AddDays 增加天数
func (receiver UTime) AddDays(days int) UTime {
	return UTime{carbon.Now().AddDays(days).Carbon2Time()}
}

// AddHour 增加1小时
func (receiver UTime) AddHour() UTime {
	return UTime{carbon.Now().AddHour().Carbon2Time()}
}

// AddHours 增加小时数
func (receiver UTime) AddHours(hours int) UTime {
	return UTime{carbon.Now().AddHours(hours).Carbon2Time()}
}

// AddMinute 增加1分钟
func (receiver UTime) AddMinute() UTime {
	return UTime{carbon.Now().AddMinute().Carbon2Time()}
}

// AddMinutes 增加分钟数
func (receiver UTime) AddMinutes(minutes int) UTime {
	return UTime{carbon.Now().AddMinutes(minutes).Carbon2Time()}
}

// AddSecond 增加1秒
func (receiver UTime) AddSecond() UTime {
	return UTime{carbon.Now().AddSecond().Carbon2Time()}
}

// AddSeconds 增加秒数
func (receiver UTime) AddSeconds(seconds int) UTime {
	return UTime{carbon.Now().AddSeconds(seconds).Carbon2Time()}
}

// SubYear 减去1年
func (receiver UTime) SubYear() UTime {
	return UTime{carbon.Now().SubYear().Carbon2Time()}
}

// SubYears 减去年数
func (receiver UTime) SubYears(years int) UTime {
	return UTime{carbon.Now().SubYears(years).Carbon2Time()}
}

// SubMonth 减去1月
func (receiver UTime) SubMonth() UTime {
	return UTime{carbon.Now().SubMonth().Carbon2Time()}
}

// SubMonths 减去月数
func (receiver UTime) SubMonths(months int) UTime {
	return UTime{carbon.Now().SubMonths(months).Carbon2Time()}
}

// SubDay 减去1天
func (receiver UTime) SubDay() UTime {
	return UTime{carbon.Now().SubDay().Carbon2Time()}
}

// SubDays 减去天数
func (receiver UTime) SubDays(days int) UTime {
	return UTime{carbon.Now().SubDays(days).Carbon2Time()}
}

// SubHour 减去1小时
func (receiver UTime) SubHour() UTime {
	return UTime{carbon.Now().SubHour().Carbon2Time()}
}

// SubHours 减去小时数
func (receiver UTime) SubHours(hours int) UTime {
	return UTime{carbon.Now().SubHours(hours).Carbon2Time()}
}

// SubMinute 减去1分钟
func (receiver UTime) SubMinute() UTime {
	return UTime{carbon.Now().SubMinute().Carbon2Time()}
}

// SubMinutes 减去分钟数
func (receiver UTime) SubMinutes(minutes int) UTime {
	return UTime{carbon.Now().SubMinutes(minutes).Carbon2Time()}
}

// SubSecond 减去1秒钟
func (receiver UTime) SubSecond() UTime {
	return UTime{carbon.Now().SubSecond().Carbon2Time()}
}

// SubSeconds 减去秒钟数
func (receiver UTime) SubSeconds(seconds int) UTime {
	return UTime{carbon.Now().SubSeconds(seconds).Carbon2Time()}
}

// ToTime UTime对象转time.Time对象
func (receiver UTime) ToTime() time.Time {
	return receiver.Time
}
