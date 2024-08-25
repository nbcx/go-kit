package tz

import (
	"time"
)

const (
	// FullFormat 最常用的格式
	FullFormat = "2006-01-02 15:04:05"
	// DayFormat DayFormat
	DayFormat = "2006-01-02"
	// HourFormat HourFormat
	HourFormat = "2006-01-02-15"
)

// 当前时区
var timeZoneLocation = time.Local
var timeDiff int64

// GetTodayStr GetTodayStr
func GetTodayStr() string {
	return GetLocalStr(Now().UTC())
}

// 获取格式化时间
// eg：
// format "2006-01-02 15:04:05" or "2006-01-02" ...
func GetFormatStr(format string) string {
	return UTCToLocal(Now().UTC()).Format(format)
}

// ChinaTimezone ChinaTimezone
func ChinaTimezone() *time.Location {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return loc
}

// GetLocalStr change utc time to local date str
func GetLocalStr(base time.Time) string {
	return base.In(GetTimezone()).Format(DayFormat)
}

// UTCToLocal 如果想要将本地时间转换成UTC，直接用UTC()方法即可
// 如果解析字符串，对应的是本地时间且字符串中没有时区，使用time.ParseInLocation(ChinaTimeZone())
func UTCToLocal(base time.Time) time.Time {
	return base.In(GetTimezone())
}

// IsSameDay check if two time is same day locally
func IsSameDay(l time.Time, r time.Time) bool {
	return GetLocalStr(l) == GetLocalStr(r)
}

// GetNowTsMs GetNowTsMs
func GetNowTsMs() int64 { //nolint
	return Now().UnixNano() / int64(time.Millisecond)
}

// GetNowTs GetNowTs
func GetNowTs() int64 { //nolint
	return Now().Unix()
}

// GetDays 获取两个时间的自然天数
func GetDays(to, from time.Time) int64 {
	start, now := GetZeroTime(to), GetZeroTime(from)
	day := int64(now.Sub(start).Hours()) / 24
	if day < 0 {
		return -day
	}
	return day
}

// Schedule Schedule
func Schedule(what func(), delay time.Duration, stop chan bool) {
	DynamicSchedule(what, &delay, stop)
}

// LocalNow LocalNow
func LocalNow() time.Time {
	return UTCToLocal(Now().UTC())
}

// Local Local
func Local(msec int64) time.Time {
	return UTCToLocal(time.UnixMilli(msec).UTC())
}

// Parse Parse
// eg: Parse("2006-01-02 15:04:05", "2022-12-16 12:04:10")
func Parse(layout, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, GetTimezone())
}

// Str
func Str(t time.Time) string {
	return t.Format(FullFormat)
}

// DynamicSchedule 可以动态修改延迟时间、可关闭的定时器
func DynamicSchedule(what func(), delayAddr *time.Duration, stop chan bool) {
	go func() {
		for {
			select {
			case <-time.After(*delayAddr):
				what()
			case <-stop:
				return
			}
		}
	}()
}

// GetZeroTime 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// GetFirstDateOfMonth 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// SetTimeZone 设置时区
func SetTimeZone(timeZone string) error {
	// log.Info("set time zone: %s", timeZone)
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return err
	}

	timeZoneLocation = loc
	return nil
}

// GetTimezone 获取时区
func GetTimezone() *time.Location {
	return timeZoneLocation
}

// GetLastWeek 获取指定多少周的信息
// n < 0 为过去多少周
// n = 0 为本周
// n > 0 为之后多少周
func GetWeek(t time.Time, n int) (year, week int) {
	return GetWeekStartTime(t, n).ISOWeek()
}

// GetLastWeek 获取指定多少周的开始时间, 及周一的零点
// n < 0 为过去多少周
// n = 0 为本周
// n > 0 为之后多少周
func GetWeekStartTime(t time.Time, n int) time.Time {
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}

	tt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, GetTimezone()).
		AddDate(0, 0, offset).AddDate(0, 0, 7*n)

	return tt
}

// 获取当前时间
func Now() time.Time {
	if timeDiff == 0 {
		return time.Now()
	}
	return time.UnixMilli(time.Now().UnixMilli() - timeDiff)
}

// 设置系统当前时间
// 此处设置的系统时间，只正对tz下的相关函数起作用
func SetSystemTime(dt string, timeZones ...string) error {
	if dt == "" {
		timeDiff = 0
		return nil
	}
	for _, v := range timeZones {
		if err := SetTimeZone(v); err != nil {
			return err
		}
	}
	t, err := Parse(FullFormat, dt)
	if err != nil {
		return err
	}
	// log.Info("set system time: %s", dt)
	timeDiff = time.Now().UnixMilli() - t.UnixMilli()
	return nil
}
